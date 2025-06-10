package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Jeet-baldha/codelink-backend/internal/config"
	"github.com/Jeet-baldha/codelink-backend/internal/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func SetupRoutes() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		filter := bson.M{"email": "jeetbaldha56@gmail.com"}

		var user model.User
		err := config.UserCollection.FindOne(context.TODO(), filter).Decode(&user)
		if err != nil {
			log.Println("User not found:", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)

	})

	return router

}
