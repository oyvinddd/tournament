package app

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
	"tournament/handler"
	"tournament/tournament"
	"tournament/user"
)

const (
	signInPath string = "/api/v1/sign_in"

	createTournamentPath string = "/api/v1/tournaments"

	joinTournamentPath string = "/api/v1/tournament"

	invitationsPath string = "/api/v1/tournaments/invitations"

	getTournamentPath string = "/api/v1/tournaments"
)

const (
	serverReadTimeout = time.Second * 15

	serverWriteTimeout = time.Second * 15
)

type App struct {
	server http.Server
}

func New(address string) *App {

	//dbURI := buildDatabseURI("localhost:5432/postgres", "postgres", "Test1234")
	dbURI := "postgresql://postgres:Test1234@localhost:5432/postgres"
	fmt.Printf("PG connection URL: %v\n", dbURI)
	dbConn, err := pgxpool.New(context.Background(), dbURI)
	_, err = dbConn.Exec(context.Background(), ";")
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatalf("Error connecting to database: %f\n", err)
	}

	authService := user.NewService(dbConn)
	tournamentService := tournament.NewService(dbConn)

	userHandler := handler.NewUserHandler(authService)
	tournamentHandler := handler.NewTournamentHandler(tournamentService)

	router := httprouter.New()
	router.POST(signInPath, userHandler.SignIn)
	router.POST(createTournamentPath, handler.AuthMiddleware(tournamentHandler.Create))
	router.GET(getTournamentPath, handler.AuthMiddleware(tournamentHandler.Get))
	router.PUT(joinTournamentPath, handler.AuthMiddleware(tournamentHandler.Join))
	router.POST(invitationsPath, handler.AuthMiddleware(userHandler.InviteUser))

	return &App{server: http.Server{
		Addr:         address,
		Handler:      router,
		ReadTimeout:  serverReadTimeout,
		WriteTimeout: serverWriteTimeout,
	}}
}

func (app *App) Run() error {
	return app.server.ListenAndServe()
}

func buildDatabseURI(host, username, password string) string {
	return fmt.Sprintf("postgresql://%s:%s@%s", username, password, host)
}
