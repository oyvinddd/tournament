package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	Service interface {

		// SignIn using Apple
		SignIn(ctx context.Context, identityToken string) (*SignInContainer, error)

		// InviteUser to the inviting users tournament
		InviteUser(ctx context.Context, userID string) error

		// GetInvitations to tournaments for a given user
		GetInvitations(ctx context.Context) ([]Invite, error)
	}

	liveService struct {
		connection *pgxpool.Pool
	}
)

func NewService() Service {
	return &liveService{}
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
	if err := service.connection.QueryRow(ctx, query, id, email, username).Scan(&id, &email, &username); err != nil {
		return nil, err
	}

	user := New(email, username)
	user.ID = id

	// TODO: generate access token for the newly signed in user

	return NewSignInContainer(*user, ""), nil
}

func (service *liveService) InviteUser(ctx context.Context, userID string) error {
	_, err := uuid.FromBytes([]byte(userID))
	if err != nil {
		return err
	}
	return nil
}

func (service *liveService) GetInvitations(ctx context.Context) ([]Invite, error) {
	return nil, nil
}
