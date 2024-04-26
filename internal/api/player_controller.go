package api

import (
	"encoding/json"
	"fmt"
	"github.com/michael-martinez-dev/globalwarfront-server/internal/db"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/michael-martinez-dev/globalwarfront-server/internal/models"
	"github.com/michael-martinez-dev/globalwarfront-server/internal/services/player"
)

type PlayerController interface {
	Register(w http.ResponseWriter, r *http.Request)
}

type playerController struct {
	service player.PlayerService
}

func NewPlayerController() PlayerController {
	database := db.NewPostgresDB("localhost", 5432, "devuser", "devpassword", "global_warfront")
	err := database.Connect()
	if err != nil {
		log.Panic(err)
	}
	defer database.Close()

	if err := database.Healthcheck(); err != nil {
		log.Panic(err)
	}
	repo := player.NewPlayerRepo(database)
	playerService := player.NewPlayerService(repo)
	return &playerController{
		service: playerService,
	}
}

func (controller *playerController) Register(w http.ResponseWriter, r *http.Request) {
	var p models.Player

	log.Infof("Processing new request")
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		log.Error(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if !p.Validate() {
		http.Error(w, "Invalid values in body", http.StatusBadRequest)
		return
	}

	if err := controller.service.Register(p); err != nil {
		log.Error(err)
		http.Error(w, "Failed to register player", http.StatusInternalServerError)
		return
	}

	_, err := fmt.Fprintln(w, "Successfully registered player")
	if err != nil {
		log.Error(err)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
