package auth

import (
	"os"

	"github.com/golang-jwt/jwt"
)

type Service interface {
	GenerateToken(userId uint) (string, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

var SECRET_KEY = []byte(os.Getenv("SECRET_KEY"))

func (s *jwtService) GenerateToken(userId uint) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}
