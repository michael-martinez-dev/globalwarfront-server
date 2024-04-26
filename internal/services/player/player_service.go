package player

import (
	"github.com/michael-martinez-dev/globalwarfront-server/internal/models"
)

type PlayerService interface {
	Register(player models.Player) error
}

type playerService struct {
	repo PlayerRepo
}

func NewPlayerService(repo PlayerRepo) PlayerService {
	return &playerService{repo}
}

func (service *playerService) Register(player models.Player) error {
	return service.repo.RegisterPlayer(player)
}
