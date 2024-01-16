package handler

import (
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
	usr, err := handler.service.SignIn(r.Context())

	if err != nil {
		respondWithStatus(w, http.StatusUnauthorized)
		return
	}

	respondWithJSON(w, usr)
}

func (handler AuthHandler) InviteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
