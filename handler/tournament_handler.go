package handler

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"tournament/tournament"
)

type TournamentHandler struct {
	service tournament.Service
}

func NewTournamentHandler(service tournament.Service) *TournamentHandler {
	return &TournamentHandler{service: service}
}

func (handler TournamentHandler) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// get the user ID from the current request context
	userID, err := userIDFromContext(r.Context())
	if err != nil {
		respondWithStatus(w, http.StatusBadRequest)
		return
	}

	var requestBody struct {
		Title string `json:"title"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		respondWithStatus(w, http.StatusBadRequest)
		return
	}

	trn, err := handler.service.CreateTournament(r.Context(), userID, requestBody.Title)
	if err != nil {
		respondWithStatus(w, http.StatusBadRequest)
		return
	}

	respondWithJSON(w, trn)
}

func (handler TournamentHandler) Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// get the user ID from the current request context
	userID, err := userIDFromContext(r.Context())
	if err != nil {
		respondWithStatus(w, http.StatusBadRequest)
		return
	}

	// get the data for the tournament the current user is a part of
	trn, err := handler.service.GetTournament(r.Context(), userID)
	if err != nil {
		respondWithStatus(w, http.StatusNotFound)
		return
	}

	respondWithJSON(w, trn)
}

func (handler TournamentHandler) Join(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// get the user ID from the current request context
	userID, err := userIDFromContext(r.Context())
	if err != nil {
		respondWithStatus(w, http.StatusBadRequest)
		return
	}

	var joinRequest tournament.JoinRequest
	if err := json.NewDecoder(r.Body).Decode(&joinRequest); err != nil {
		respondWithStatus(w, http.StatusBadRequest)
		return
	}

	err = handler.service.JoinTournament(r.Context(), userID, joinRequest.TournamentID)
	if err != nil {
		respondWithStatus(w, http.StatusUnauthorized)
		return
	}

	respondWithStatus(w, http.StatusOK)
}
