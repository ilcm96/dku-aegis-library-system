package repository

import (
	"context"
	"log"

	"github.com/ilcm96/dku-aegis-library/ent"
	"github.com/ilcm96/dku-aegis-library/ent/user"
	"github.com/ilcm96/dku-aegis-library/model"
)

type UserRepository interface {
	Create(user *model.User) error
	FindUserById(id int) (*ent.User, error)
}

type userRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) UserRepository {
	return &userRepository{
		client: client,
	}
}

func (ur *userRepository) Create(user *model.User) error {
	_, err := ur.client.User.Create().
		SetID(user.Id).
		SetPassword(user.Password).
		SetName(user.Name).
		Save(context.Background())

	if err != nil {
		log.Println("ERR: save to db failed |", err)
	}
	return err
}

func (ur *userRepository) FindUserById(id int) (*ent.User, error) {
	return ur.client.User.Query().
		Where(user.ID(id)).
		First(context.Background())
}
