package usercredential

import (
	"context"

	dbmodel "github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/db/model"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type Manager struct {
}

func NewManager() *Manager {

}

func (m *Manager) HashByUsername(ctx context.Context, username string) ([]byte, error) {
	u, err := dbmodel.Users(
		dbmodel.UserWhere.Username.EQ(username),
		qm.Load(
			dbmodel.UserRels.IDUserCredential,
		),
	).One(ctx, nil)
	if err != nil {
		return nil, err
	}

	return u.R.IDUserCredential.Hash, nil
}
