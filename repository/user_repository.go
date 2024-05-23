package repository

import (
	"context"
	"github.com/ilcm96/dku-aegis-library/ent"
	"github.com/ilcm96/dku-aegis-library/ent/user"
	"github.com/ilcm96/dku-aegis-library/model"
)

type UserRepository interface {
	Create(user *model.User) error
	Update(user *model.User) error
	FindUserById(id int) (*ent.User, error)
	Withdraw(id int) error
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

	return err
}

func (ur *userRepository) Update(user *model.User) error {
	_, err := ur.client.User.UpdateOneID(user.Id).
		SetPassword(user.Password).
		SetName(user.Name).
		Save(context.Background())

	return err
}

func (ur *userRepository) FindUserById(id int) (*ent.User, error) {
	return ur.client.User.Query().
		Where(user.ID(id)).
		First(context.Background())
}

func (ur *userRepository) Withdraw(id int) error {
	_, err := ur.client.User.Update().
		Where(user.ID(id)).
		SetStatus(user.StatusWITHDRAW).
		SetName("WITHDRAW_USER").
		SetPassword("WITHDRAW_USER").
		Save(context.Background())

	return err
}
