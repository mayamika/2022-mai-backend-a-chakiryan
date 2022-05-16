package auth

import (
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/ent"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/model/token"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) tokenFromUser(u *ent.User) (string, error) {
	t := token.Token{
		UserID: u.ID,
	}
	return t.SignedString()
}
