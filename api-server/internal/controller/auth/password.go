package auth

import "golang.org/x/crypto/bcrypt"

func (c *Controller) compareHashAndPassword(hashedPassword, password string) (ok bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return
	}
	return true
}

func (c *Controller) passwordHash(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
