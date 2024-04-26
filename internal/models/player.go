package models

import (
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
	safeUser, _ := regexp.MatchString(
		`^\w+$`, player.Username)
	safeEmail, _ := regexp.MatchString(
		`^.*@.*\.\w+$`, player.Email)
	safePassword, _ := regexp.MatchString(
		`^.{8,}$`, player.Password)
	log.Infof("Validation results: %v %v", safeUser, safePassword)
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
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(player.Password))
	return err == nil
}
