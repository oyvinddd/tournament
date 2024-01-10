package tournament

import (
	"github.com/google/uuid"
	"time"
	"tournament/user"
)

type (
	Tournament struct {

		// ID of the tournament
		ID uuid.UUID `json:"id"`

		// Admin the user administrating said tournament (FK)
		AdminID uuid.UUID `json:"admin_id"`

		// Title of the tournament
		Title string `json:"title"`

		// Code a six-digit code used whenever someone wants to join a tournament
		Code int `json:"-"`

		// Scoreboard shows the current status if the tournament
		Scoreboard []user.User `json:"scoreboard"`

		// LatestMatches is a list of the latest matches players in the tournament
		LatestMatches []Match `json:"latest_matches"`
	}

	Match struct {
		ID uuid.UUID `id:"id"`

		Winner string `json:"winner"`

		Loser string `json:"loser"`

		Date time.Time `json:"date"`
	}

	JoinRequest struct {
		ID uuid.UUID

		Code int
	}
)

func New(title string, code int) *Tournament {
	return &Tournament{
		ID:    uuid.New(),
		Title: title,
		Code:  code,
	}
}
