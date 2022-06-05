package token

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

type Token struct {
	UserID int
}

func (t Token) SignedString() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims{
		ID: t.UserID,
	})
	return token.SignedString(hmacSecret)
}

func FromSignedString(tokenString string) (Token, error) {
	var c claims
	_, err := jwt.ParseWithClaims(tokenString, &c, keyFunc)
	if err != nil {
		return Token{}, err
	}

	return Token{
		UserID: c.ID,
	}, nil
}

type claims struct {
	ID int
}

func (c claims) Valid() error {
	return nil
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	return hmacSecret, nil
}
