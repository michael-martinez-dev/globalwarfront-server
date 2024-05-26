package player

import (
	"database/sql"
	log "github.com/sirupsen/logrus"

	"github.com/michael-martinez-dev/globalwarfront-server/internal/db"
	"github.com/michael-martinez-dev/globalwarfront-server/internal/models"
)

type PlayerRepo interface {
	registerPlayer(player models.Player) error
	retrievePlayerByUsername(username string) (*models.Player, error)
}

type playerRepo struct {
	DB db.DB
}

func NewPlayerRepo(db db.DB) PlayerRepo {
	return &playerRepo{db}
}

func (repo *playerRepo) registerPlayer(player models.Player) error {
	log.Info("Adding player to database...")
	log.Infof("%s", player.PrintStr())
	if err := repo.DB.Healthcheck(); err != nil {
		log.Error(err)
		return err
	}
	query := "INSERT INTO players(email, username, password_hash) VALUES ($1, $2, $3)"
	stmt, err := repo.DB.GetDB().Prepare(query)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Error(err)
		}
	}(stmt)

	_, err = stmt.Exec(player.Email, player.Username, player.GetPasswordHash())
	if err != nil {
		log.Error(err)
		return err
	}

	log.Infof("Successfully added new player to repo")
	return nil
}

func (repo *playerRepo) retrievePlayerByUsername(username string) (*models.Player, error) {
	log.Info("Retrieving player from database...")
	log.Infof("%s", username)
	if err := repo.DB.Healthcheck(); err != nil {
		return nil, err
	}

	var retrievedPlayer models.Player

	query := `SELECT username, password_hash FROM players WHERE username = $1`

	result := repo.DB.GetDB().QueryRow(query, username)

	if result != nil {
		if err := result.Scan(&retrievedPlayer.Username, &retrievedPlayer.Password); err != nil {
			return nil, err
		}
	}
	return &retrievedPlayer, nil
}
