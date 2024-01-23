package tournament

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	Service interface {
		CreateTournament(ctx context.Context, userID uuid.UUID, title string) (*Tournament, error)

		GetTournament(ctx context.Context, userID uuid.UUID) (*Tournament, error)

		JoinTournament(ctx context.Context, userID, tournamentID uuid.UUID) error
	}

	liveService struct {
		db *pgxpool.Pool
	}
)

func NewService(connection *pgxpool.Pool) Service {
	return &liveService{db: connection}
}

func (service *liveService) CreateTournament(ctx context.Context, userID uuid.UUID, title string) (*Tournament, error) {

	// first we insert a new row into the table representing a tournament
	query1 := `INSERT INTO tr.tournaments (title) VALUES ($1) RETURNING id;`

	var tournamentID uuid.UUID
	if err := service.db.QueryRow(ctx, query1, title).Scan(&tournamentID); err != nil {
		return nil, err
	}

	tournament := Tournament{
		ID:            tournamentID,
		Title:         title,
		ResetInterval: ResetMonthly,
	}

	// secondly we update the current user row with the newly created tournament ID
	// the current user is now the tournament owner
	query2 := `UPDATE INTO tr.users (tournament_id, tournament_role) VALUES ($1, $2) WHERE id = $3;`
	service.db.QueryRow(ctx, query2, tournamentID)

	return &tournament, nil
}

func (service *liveService) GetTournament(ctx context.Context, userID uuid.UUID) (*Tournament, error) {
	tx, err := service.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback(ctx)

	/*
		query := `SELECT id, title FROM tr.tournaments WHERE url_name = $1`
		var a artist.Artist
		row := repo.db.conn.QueryRow(ctx, stmt, urlName)
		if err := row.Scan(&a.ID, &a.Name, &a.ProfileImage, &a.BackgroundImage); err != nil {
			return nil, errArtistNotFound
		}
		return &a, nil
	*/
}

func (service *liveService) JoinTournament(ctx context.Context, userID, tournamentID uuid.UUID) error {
	query := `
UPDATE 
    tr.users 
SET 
    tournament_id = $1, tournament_role = 1, score = 0, matches_played = 0, matches_won = 0 
WHERE 
    id = $2 
  AND 
    EXISTS (SELECT 1 FROM tr.tournaments WHERE id = $1);`

	_, err := service.db.Exec(ctx, query, tournamentID, userID)
	return err
}
