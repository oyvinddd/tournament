package tournament

import (
	"github.com/google/uuid"
	"time"
	"tournament/user"
)

const (
	ResetWeekly ResetInterval = 0

	ResetMonthly ResetInterval = 1
)

type (
	ResetInterval int

	Tournament struct {

		// ID of the tournament
		ID uuid.UUID `json:"id"`

		// Title of the tournament
		Title string `json:"title"`

		ResetInterval ResetInterval `json:"reset_interval"`

		// Scoreboard shows the current status if the tournament
		Scoreboard []user.User `json:"scoreboard"`

		CreatedAt time.Time `json:"-"`
	}

	Match struct {
		ID uuid.UUID `id:"id"`

		Winner string `json:"winner"`

		Loser string `json:"loser"`

		Date time.Time `json:"date"`
	}

	JoinRequest struct {
		TournamentID uuid.UUID `json:"tournament_id"`
	}
)

func New(title string) *Tournament {
	return &Tournament{
		ID:    uuid.New(),
		Title: title,
	}
}
