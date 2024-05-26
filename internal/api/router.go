package api

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
	"strings"
)

func NewRouter() http.Handler {
	playerHandler := NewPlayerController()
	router := mux.NewRouter()

	router.HandleFunc("/players/register", playerHandler.Register).Methods("POST")
	router.HandleFunc("/players/login", playerHandler.Login).Methods("Post")
	// TODO: /login with jwt

	c := cors.New(cors.Options{
		AllowOriginFunc:  func(origin string) bool { return strings.HasPrefix(origin, "http://localhost:") },
		AllowedMethods:   []string{"POST", "GET", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "X-Requested-With", "Authorization"},
		AllowCredentials: true,
		Debug:            true,
	})

	return c.Handler(router)
}
