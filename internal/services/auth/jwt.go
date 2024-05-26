package auth

import (
	"github.com/michael-martinez-dev/globalwarfront-server/internal/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthService interface {
	CreateToken(player models.Player) (string, error)
	ValidateToken(bearerToken string) (jwt.MapClaims, bool)
}

type jwtAuthService struct {
	secret string
}

func NewJwtAuth(secret string) AuthService {
	return &jwtAuthService{
		secret: secret,
	}
}

func (auth *jwtAuthService) CreateToken(player models.Player) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"username": player.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(auth.secret))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func (auth *jwtAuthService) ValidateToken(tokenStr string) (jwt.MapClaims, bool) {
	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(auth.secret), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return claims, false
	}

	return claims, true
}
