package helpers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/souta17/go/entities"
)

var mySecret = []byte("mysecret")

type JWTclaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(user *entities.User) (string, error) {
	claims := JWTclaims{
		user.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySecret)
	return ss, err
}
