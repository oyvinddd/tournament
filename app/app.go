package app

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
	"tournament/handler"
	"tournament/tournament"
	"tournament/user"
)

const (
	signInPath string = "/api/v1/sign_in"

	createTournamentPath string = "/api/v1/tournament/new"

	joinTournamentPath string = "/api/v1/tournament"

	invitationsPath string = "/api/v1/tournament/invitations"

	getTournamentPath string = "/api/v1/tournament"
)

const (
	serverReadTimeout time.Duration = 15

	serverWriteTimeout time.Duration = 15
)

type App struct {
	server http.Server
}

func New(address string) *App {

	authService := user.NewService()
	tournamentService := tournament.NewService()

	userHandler := handler.NewUserHandler(authService)
	tournamentHandler := handler.NewTournamentHandler(tournamentService)

	router := httprouter.New()
	router.POST(signInPath, userHandler.SignIn)
	router.POST(createTournamentPath, handler.AuthMiddleware(tournamentHandler.Create))
	router.GET(getTournamentPath, tournamentHandler.Get)
	router.PUT(joinTournamentPath, tournamentHandler.Join)

	return &App{server: http.Server{
		Addr:         address,
		Handler:      router,
		TLSConfig:    nil,
		ReadTimeout:  serverReadTimeout,
		WriteTimeout: serverWriteTimeout,
	}}
}

func (app *App) Run() error {
	return app.server.ListenAndServe()
}
