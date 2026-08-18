package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userId string, expiresAt int64, secret []byte) (string, error) {
	claims := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": userId,
			"exp": expiresAt,
			"iat": time.Now().Unix(),
		},
	)

	token, err := claims.SignedString(secret)
	if err != nil {
		return "", err
	}

	return token, nil
}

func VerifyJWT(tokenString string, secret []byte) (*jwt.Token, error) {
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		},
	)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token provided")
	}

	return token, nil
}
