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

		TournamentsWon int `json:"tournaments_won"`

		LastSeen time.Time `json:"last_seen"`
	}

	Invite struct {
		TournamentID uuid.UUID `json:"tournament_id"`

		InviteeID uuid.UUID `json:"invitee_id"`

		CreatedAt time.Time `json:"created_at"`
	}

	JoinRequest struct {
		TournamentID uuid.UUID `json:"tournament_id"`
	}

	SignInContainer struct {
		User User `json:"user"`

		AccessToken string `json:"access_token"`
	}
)

func New(username, email string) *User {
	return &User{
		ID:             uuid.New(),
		Username:       username,
		Email:          email,
		CreatedAt:      time.Now(),
		TournamentRole: RoleNoRole,
		LastSeen:       time.Now(),
	}
}

func NewSignInContainer(user User, accessToken string) *SignInContainer {
	return &SignInContainer{User: user, AccessToken: accessToken}
}
