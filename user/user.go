package user

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID uuid.UUID `json:"id"`

	Username string `json:"username"`

	Email string `json:"-"`

	CreatedAt time.Time `json:"-"`

	TournamentID *uuid.UUID `json:"-"`

	Score int `json:"score"`

	LastSeen time.Time `json:"last_seen"`
}

func New(username, email string) *User {
	return &User{
		ID:        uuid.New(),
		Username:  username,
		Email:     email,
		CreatedAt: time.Now(),
	}
}
