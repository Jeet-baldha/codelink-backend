package handler

import (
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Login route hit"}`))
}
