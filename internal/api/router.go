package api

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	playerHandler := NewPlayerController()
	router := mux.NewRouter()

	router.HandleFunc("/players/register", playerHandler.Register).Methods("POST")

	return router
}
