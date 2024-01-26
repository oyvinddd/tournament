package user

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
	"tournament/jwtutil"
)

type (
	Service interface {

		// SignIn using Apple
		SignIn(ctx context.Context, identityToken string) (*SignInContainer, error)

		// InviteUser to the inviting users tournament
		InviteUser(ctx context.Context, inviterID, inviteeID uuid.UUID) error

		// GetInvitations to tournaments for a given user
		GetInvitations(ctx context.Context, userID uuid.UUID) ([]Invite, error)

		// LeaveTournament leave current tournament
		LeaveTournament(ctx context.Context, userID uuid.UUID) error

		// DeleteUser deletes current user
		DeleteUser(ctx context.Context, userID uuid.UUID) error
	}

	liveService struct {
		db *pgxpool.Pool
	}
)

func NewService(db *pgxpool.Pool) Service {
	return &liveService{db: db}
}

func (service *liveService) SignIn(ctx context.Context, identityToken string) (*SignInContainer, error) {

	// TODO: validate Apple identity token using public key from Apple

	// create query
	query := `INSERT INTO tr.users (email, username)
	VALUES ($1, $2)
	ON CONFLICT (email) DO UPDATE
	SET email = EXCLUDED.email
	RETURNING id, username, email;`

	// execute query
	var id uuid.UUID
	var email, username string
	if err := service.db.QueryRow(ctx, query, id, email, username).Scan(&id, &email, &username); err != nil {
		return nil, err
	}

	user := New(email, username)
	user.ID = id

	// TODO: generate access token for the newly signed in user
	token, err := jwtutil.GenerateToken(id.String(), 0, time.Minute*60)
	if err != nil {
		return nil, errors.New("unable to generate token")
	}

	return NewSignInContainer(*user, token), nil
}

func (service *liveService) InviteUser(ctx context.Context, inviterID, inviteeID uuid.UUID) error {

	// first find the tournament ID for the tournament of which the inviter is an admin
	query1 := `
SELECT 
    tournament_id 
FROM 
    tr.users 
WHERE id $1 AND (tournament_role = $2 OR tournament_role = $3);`

	var tournamentID *uuid.UUID
	err := service.db.QueryRow(ctx, query1, inviterID, RoleAdmin, RoleOwner).Scan(tournamentID)
	if err != nil || tournamentID == nil {
		return errors.New("tournament not found")
	}

	// we can now insert a new invitation row into the table
	query2 := `INSERT INTO tr.invitations (tournament_id, invitee_id) VALUES ($1, $2);`
	_, err = service.db.Exec(ctx, query2, tournamentID, inviteeID)
	return err
}

func (service *liveService) GetInvitations(ctx context.Context, userID uuid.UUID) ([]Invite, error) {
	query := `SELECT * FROM tr.invitations WHERE invitee_id = $1;`

	rows, err := service.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	var invitations = make([]Invite, 0)
	for rows.Next() {

		var invite Invite
		if err := rows.Scan(&invite); err != nil {
			rows.Close()
			return nil, err
		}

		invitations = append(invitations, invite)
	}
	rows.Close()
	return invitations, nil
}

func (service *liveService) LeaveTournament(ctx context.Context, userID uuid.UUID) error {
	query := `
UPDATE 
    tr.users 
SET 
    tournament_id = null, tournament_role = 0, score = 0, matches_played = 0, matches_won = 0
WHERE
    id = $1;
`
	_, err := service.db.Exec(ctx, query, userID)
	return err
}

func (service *liveService) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	_, err := service.db.Exec(ctx, `DELETE * FROM tr.users WHERE id = $1;`, userID)
	return err
}
