package app

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
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
	router.GET(createTournamentPath, tournamentHandler.Create)
	router.GET(getTournamentPath, tournamentHandler.Get)
	router.PUT(joinTournamentPath, tournamentHandler.Join)

	// TODO: https://justinas.org/writing-http-middleware-in-go
	return &App{server: http.Server{
		Addr:              address,
		Handler:           router,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}}
}

func (app *App) Run() error {
	return app.server.ListenAndServe()
}
