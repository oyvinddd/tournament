package tournament

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	Service interface {
		CreateTournament(ctx context.Context, userID uuid.UUID, title string) (*Tournament, error)

		// JoinTournament joins a given tournament
		JoinTournament(ctx context.Context, userID, tournamentID uuid.UUID) error

		GetTournament(ctx context.Context, userID uuid.UUID) (*Tournament, error)
	}

	liveService struct {
		db *pgxpool.Pool
	}
)

func NewService(db *pgxpool.Pool) Service {
	return &liveService{db: db}
}

func (service *liveService) CreateTournament(ctx context.Context, userID uuid.UUID, title string) (*Tournament, error) {
	tx, err := service.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	// first we insert a new row into the table representing a tournament
	query1 := `INSERT INTO tr.tournaments (title) VALUES ($1) RETURNING id;`

	var tournamentID uuid.UUID
	if err := tx.QueryRow(ctx, query1, title).Scan(&tournamentID); err != nil {
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
	if _, err := tx.Exec(ctx, query2, tournamentID, 3, userID); err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err

	}
	return &tournament, nil
}

func (service *liveService) GetTournament(ctx context.Context, userID uuid.UUID) (*Tournament, error) {
	query := `SELECT * FROM tr.tournaments WHERE id = (SELECT tournament_id FROM tr.users WHERE id = $1);`

	var tournament Tournament
	if err := service.db.QueryRow(ctx, query, userID).Scan(&tournament); err != nil {
		return nil, err
	}
	return &tournament, nil
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
