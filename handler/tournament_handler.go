package handler

import (
	"encoding/json"
	"github.com/google/uuid"
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
	title := ps.ByName("title")
	t, err := handler.service.Create(title)
	if err != nil {
		respondWithStatus(w, http.StatusBadRequest)
		return
	}

	respondWithJSON(w, t)
}

func (handler TournamentHandler) Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	guid, err := uuid.FromBytes([]byte(id))
	if err != nil {
		respondWithStatus(w, http.StatusBadRequest)
		return
	}

	t, err := handler.service.Get(guid)
	if err != nil {
		respondWithStatus(w, http.StatusNotFound)
		return
	}

	respondWithJSON(w, t)
}

func (handler TournamentHandler) Join(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var req tournament.JoinRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithStatus(w, http.StatusBadRequest)
		return
	}

	t, err := handler.service.Join(req)
	if err != nil {
		respondWithStatus(w, http.StatusUnauthorized)
		return
	}

	respondWithJSON(w, t)
}
