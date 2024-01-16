package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	Service interface {

		// SignIn using Apple
		SignIn(ctx context.Context) (User, error)

		CreateUser(ctx context.Context, user User) (User, error)

		// InviteUser to the inviting users tournament
		InviteUser(ctx context.Context, userID string) error

		// GetInvitations to tournaments for a given user
		GetInvitations(ctx context.Context) ([]Invite, error)
	}

	LiveService struct {
		connection *pgxpool.Pool
	}
)

func NewService() Service {
	return &LiveService{}
}

func (service *LiveService) SignIn(ctx context.Context) (User, error) {
	return User{}, nil
}

func (service *LiveService) CreateUser(ctx context.Context, user User) (User, error) {
	query := `INSERT INTO tr.users (id, email, username) VALUES ($1, $2, $3) RETURNING (id, email, username)`

	// FIXME: returned values are never used
	var id uuid.UUID
	var email, username string
	err := service.connection.QueryRow(ctx, query, user.ID, user.Email, user.Username).Scan(&id, &email, &username)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (service *LiveService) InviteUser(ctx context.Context, userID string) error {
	_, err := uuid.FromBytes([]byte(userID))
	if err != nil {
		return err
	}
	return nil
}

func (service *LiveService) GetInvitations(ctx context.Context) ([]Invite, error) {
	return nil, nil
}
