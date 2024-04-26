package player

import (
	"database/sql"
	log "github.com/sirupsen/logrus"

	"github.com/michael-martinez-dev/globalwarfront-server/internal/db"
	"github.com/michael-martinez-dev/globalwarfront-server/internal/models"
)

type PlayerRepo interface {
	RegisterPlayer(player models.Player) error
}

type playerRepo struct {
	DB db.DB
}

func NewPlayerRepo(db db.DB) PlayerRepo {
	return &playerRepo{db}
}

func (repo *playerRepo) RegisterPlayer(player models.Player) error {
	log.Info("Adding player to repo...")

	stmt, err := repo.DB.GetDB().Prepare("INSERT INTO players(email, username, password_hash) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Error(err)
		}
	}(stmt)

	_, err = stmt.Exec(player.Email, player.Username, string(player.GetPasswordHash()))
	if err != nil {
		log.Error(err)
		return err
	}

	log.Infof("Successfully added new player to repo")
	return nil
}
