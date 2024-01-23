package handler

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"tournament/jwtutil"
)

const (
	userIDContextKey string = "user_id"
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
		next(w, r.WithContext(contextWithUserID(r.Context(), sub)), ps)
	}
}

func contextWithUserID(ctx context.Context, userID interface{}) context.Context {
	return context.WithValue(ctx, userIDContextKey, userID)
}

func userIDFromContext(ctx context.Context) (uuid.UUID, error) {
	userID, ok := ctx.Value(userIDContextKey).(string)
	if !ok {
		return uuid.Nil, errors.New("no user ID found")
	}

	guid, err := uuid.FromBytes([]byte(userID))
	if err != nil {
		return uuid.Nil, err
	}

	return guid, nil
}
