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
		TournamentID uuid.UUID `json:"tournament_id"`
	}
)

func New(title string) *Tournament {
	return &Tournament{
		ID:    uuid.New(),
		Title: title,
	}
}
