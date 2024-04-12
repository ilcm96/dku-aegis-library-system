package service

import (
	"errors"
	"log"

	"github.com/ilcm96/dku-aegis-library/model"
	"github.com/ilcm96/dku-aegis-library/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	SignUp(user *model.User) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (us *userService) SignUp(user *model.User) error {
	// 이미 가입된 아이디가 있는지 확인한다
	_, err := us.userRepo.FindUserById(user.Id)
	if err == nil { // 가입된 아이디가 있다면 err == nil 이다 (없다면 *NotFoundError 가 발생한다)
		log.Println("ERR: id already exists")
		return errors.New("id already exists")
	}

	// 비밀번호 암호화
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("ERR: failed to bcrypt", err)
		return err
	}
	user.Password = string(hashedPassword)

	return us.userRepo.Create(user)
}
