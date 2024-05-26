package api

import (
	"encoding/json"
	"fmt"
	"github.com/michael-martinez-dev/globalwarfront-server/internal/db"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/michael-martinez-dev/globalwarfront-server/internal/models"
	"github.com/michael-martinez-dev/globalwarfront-server/internal/services/auth"
	"github.com/michael-martinez-dev/globalwarfront-server/internal/services/player"
)

type PlayerController interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type playerController struct {
	service player.PlayerService
}

func NewPlayerController() PlayerController {
	// TODO: make these are env vars or use property manager
	database := db.NewPostgresDB(
		"localhost",
		5432,
		"devuser",
		"devpassword",
		"global_warfront",
	)

	err := database.Connect()
	if err != nil {
		log.Panic(err)
	}

	if err := database.Healthcheck(); err != nil {
		log.Panic(err)
	}

	repo := player.NewPlayerRepo(database)
	authService := auth.NewJwtAuth("ChangeMe") // TODO: make secret env var
	playerService := player.NewPlayerService(repo, authService)
	return &playerController{
		service: playerService,
	}
}

func (controller *playerController) Register(w http.ResponseWriter, r *http.Request) {
	var p models.Player

	log.Infof("Processing new Register request")
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

func (controller *playerController) Login(w http.ResponseWriter, r *http.Request) {
	var p models.Player

	log.Info("Processing new Login request")
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		log.Error(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if !p.Validate() {
		http.Error(w, "Invalid values in body", http.StatusBadRequest)
		return
	}

	token, err := controller.service.Login(p)
	if err != nil {
		log.Error(err)
		http.Error(w, "Error logging in player", http.StatusInternalServerError)
		return
	}

	_, err = fmt.Fprintf(w, "%s", token)
	if err != nil {
		log.Error(err)
		http.Error(w, "Error loggin in player", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
