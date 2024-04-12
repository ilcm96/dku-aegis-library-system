package repository

import (
	"github.com/ilcm96/dku-aegis-library/ent"
	"github.com/ilcm96/dku-aegis-library/model"
)

type UserRepository interface {
	Create(user *model.User) error
	FindUserById(id int) (*ent.User, error)
}
