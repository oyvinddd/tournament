package handler

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const headerKeyAuthorization string = "Authorization"

func AuthMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		token := r.Header.Get(headerKeyAuthorization)
		if token == "" {
			respondWithStatus(w, http.StatusUnauthorized)
			return
		}
		next(w, r, ps)
	}
}
