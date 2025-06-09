package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {

	router := mux.NewRouter()

	// Define the routes
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//write response hello
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Welcome to Codelink"}`))

	})

	return router

}
