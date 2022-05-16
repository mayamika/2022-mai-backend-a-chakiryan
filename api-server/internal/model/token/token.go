package token

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

type Token struct {
	UserID int
}

func (t Token) SignedString() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": t.UserID,
	})
	return token.SignedString(hmacSecret)
}

func FromSignedString(tokenString string) (Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSecret, nil
	})

	if err != nil {
		return Token{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return Token{}, fmt.Errorf("can't unpack claims: %w", err)
	}

	v := claims["id"]
	id, ok := v.(int)
	if !ok {
		return Token{}, errors.New("invalid id type")
	}

	return Token{
		UserID: id,
	}, nil
}
