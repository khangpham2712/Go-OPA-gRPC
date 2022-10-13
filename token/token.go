package token

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	jwt.StandardClaims
	Username string
	Role     string
}

func Generate(username string, role string, secretKey string) (string, error) {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(48 * time.Hour).Unix(),
		},
		Username: username,
		Role:     role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func Verify(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("dummy"), nil
		})

	if err != nil {
		return nil, err
	}
	userClaims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, errors.New("verifying process error")
	}
	if userClaims.ExpiresAt < time.Now().Local().Unix() {
		return nil, err
	}
	return userClaims, nil
}
