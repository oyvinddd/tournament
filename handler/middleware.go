package handler

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"tournament/jwtutil"
)

func AuthMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// for regular secure routes, the "typ" claim can safely be ignored
		sub, _, err := jwtutil.ValidateTokenFromRequest(r)
		if err != nil {
			//handler.RespondWithJSON(w, http.StatusUnauthorized, err.Error())
			respondWithError(w, http.StatusUnauthorized, err)
			return
		}

		// the "sub" claim should contain the user ID
		ctx := context.WithValue(r.Context(), "sub", sub)
		next(w, r.WithContext(ctx), ps)
	}
}
