package models

import (
	"fmt"
	"regexp"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type Player struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (player *Player) Validate() bool {
	log.Infof("Validating player...")

	safeEmail := true
	if player.Email != "" {
		safeEmail, _ = regexp.MatchString(
			`^.*$`, player.Email) // TODO: make better
	}

	safeUser, _ := regexp.MatchString(
		`^\w+$`, player.Username)
	safePassword, _ := regexp.MatchString(
		`^.{8,}$`, player.Password)

	return safeUser && safePassword && safeEmail
}

func (player *Player) GetPasswordHash() string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(player.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error(err)
	}
	return string(bytes)
}

func (player *Player) CheckPasswordHash(hash string) bool {
	log.Info("Checking password hashes...")
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(player.Password))
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}

func (player *Player) PrintStr() string {
	return fmt.Sprintf("{\n\tusername: %s,\n\temail: %s,\n\tpassword: REDACTED,\n}", player.Username, player.Email)
}
