package user

import (
	"github.com/google/uuid"
	"time"
)

const (
	RoleNoRole TournamentRole = 0

	RoleMember TournamentRole = 1

	RoleAdmin TournamentRole = 2

	RoleOwner TournamentRole = 3
)

type (
	TournamentRole uint

	User struct {
		ID uuid.UUID `json:"id"`

		Email string `json:"-"`

		Username string `json:"username"`

		CreatedAt time.Time `json:"-"`

		TournamentID *uuid.UUID `json:"-"`

		TournamentRole TournamentRole `json:"role"`

		Score int `json:"score"`

		MatchesPlayed int `json:"matches_played"`

		MatchesWon int `json:"matches_won"`

		LastSeen time.Time `json:"last_seen"`
	}

	Invite struct {
		TournamentID uuid.UUID `json:"tournament_id"`

		InviteeID uuid.UUID `json:"invitee_id"`

		CreatedAt time.Time `json:"created_at"`

		ExpiresAt time.Time `json:"expires_at"`
	}

	InviteRequest struct {
		TournamentID uuid.UUID `json:"tournament_id"`

		InviteeID uuid.UUID `json:"invitee_id"`
	}
)

func New(username, email string) *User {
	return &User{
		ID:             uuid.New(),
		Username:       username,
		Email:          email,
		CreatedAt:      time.Now(),
		TournamentID:   nil,
		TournamentRole: RoleNoRole,
		Score:          0,
		MatchesPlayed:  0,
		MatchesWon:     0,
		LastSeen:       time.Now(),
	}
}
