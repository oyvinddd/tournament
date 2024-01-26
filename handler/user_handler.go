package handler

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"tournament/user"
)

type AuthHandler struct {
	service user.Service
}

func NewUserHandler(service user.Service) *AuthHandler {
	return &AuthHandler{service: service}
}

func (handler AuthHandler) SignIn(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	token := struct {
		IdentityToken string `json:"identity_token"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&token); err != nil {
		respondWithStatus(w, http.StatusBadRequest)
		return
	}

	container, err := handler.service.SignIn(r.Context(), token.IdentityToken)

	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err)
		//respondWithStatus(w, http.StatusUnauthorized)
		return
	}

	respondWithJSON(w, container)
}

func (handler AuthHandler) InviteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
