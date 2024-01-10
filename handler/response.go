package handler

import (
	"encoding/json"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&payload)
}

func respondWithStatus(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
	_, _ = w.Write([]byte(nil))
}

func respondWithError(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	// TODO: return a domain specific error code in the response body as well as the error message
	body := struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}{Code: 1, Message: err.Error()}
	_ = json.NewEncoder(w).Encode(body)
}
