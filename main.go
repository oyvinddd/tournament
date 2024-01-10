package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"tournament/handler"
	"tournament/tournament"
	"tournament/user"
)

const (
	signInPath           string = "/api/v1/sign_in"
	createTournamentPath string = "/api/v1/tournaments"
	getTournamentPath    string = "/api/v1/tournaments/:id"
)

func main() {

	authService := user.NewService()
	tournamentService := tournament.NewService()

	userHandler := handler.NewUserHandler(authService)
	tournamentHandler := handler.NewTournamentHandler(tournamentService)

	router := httprouter.New()
	router.POST(signInPath, userHandler.SignIn)
	router.GET(createTournamentPath, tournamentHandler.Create)
	router.GET(getTournamentPath, tournamentHandler.Get)
	router.PUT(getTournamentPath, tournamentHandler.Join)

	log.Fatal(http.ListenAndServe(":8080", router))
}
