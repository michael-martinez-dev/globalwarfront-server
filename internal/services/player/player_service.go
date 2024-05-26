package player

import (
	"errors"
	"github.com/michael-martinez-dev/globalwarfront-server/internal/models"
	"github.com/michael-martinez-dev/globalwarfront-server/internal/services/auth"
)

type PlayerService interface {
	Register(player models.Player) error
	Login(player models.Player) (string, error)
}

type playerService struct {
	repo PlayerRepo
	auth auth.AuthService
}

func NewPlayerService(repo PlayerRepo, authService auth.AuthService) PlayerService {
	return &playerService{
		repo: repo,
		auth: authService,
	}
}

func (service *playerService) Register(player models.Player) error {
	return service.repo.registerPlayer(player)
}

func (service *playerService) Login(player models.Player) (string, error) {
	retrievedPlayer, err := service.repo.retrievePlayerByUsername(player.Username)
	if err != nil {
		return "", err
	}

	if !player.CheckPasswordHash(retrievedPlayer.Password) {
		return "", errors.New("player passwords did not match")
	}

	token, err := service.auth.CreateToken(player)
	if err != nil {
		return "", err
	}

	return token, nil
}
