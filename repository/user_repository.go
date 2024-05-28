package repository

import (
	"context"
	"github.com/ilcm96/dku-aegis-library/ent"
	"github.com/ilcm96/dku-aegis-library/ent/user"
	user2 "github.com/ilcm96/dku-aegis-library/ent/user"
	"github.com/ilcm96/dku-aegis-library/model"
)

type UserRepository interface {
	Create(user *model.User) error
	CreateWithdrawUser(user *model.User) error
	FindAllUser() ([]*ent.User, error)
	FindUserById(id int) (*ent.User, error)
	Withdraw(id int) error
	ChangeStatus(id int, status user.Status) error
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

func (ur *userRepository) CreateWithdrawUser(user *model.User) error {
	_, err := ur.client.User.UpdateOneID(user.Id).
		SetPassword(user.Password).
		SetName(user.Name).
		SetStatus(user2.StatusPENDING).
		Save(context.Background())

	return err
}

func (ur *userRepository) FindAllUser() ([]*ent.User, error) {
	return ur.client.User.Query().
		All(context.Background())
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

func (ur *userRepository) ChangeStatus(id int, status user.Status) error {
	return ur.client.User.UpdateOneID(id).
		SetStatus(status).
		Exec(context.Background())
}
