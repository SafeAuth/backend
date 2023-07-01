package jwt

import (
	"errors"

	jwt "github.com/golang-jwt/jwt/v4"
)

func Validar(tokens string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokens, func(token *jwt.Token) (any, error) {
		// we verify if the cryptography method matches
		if _, err := token.Method.(*jwt.SigningMethodHMAC); !err {
			return nil, errors.New("invalid algorithm")
		}
		return nil, nil
	})

	// if there is an error, we return it
	if err != nil {
		return nil, err
	}

	// if the token is valid, we return the claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
