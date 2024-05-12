package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ilcm96/dku-aegis-library/ent"
	"time"

	"github.com/ilcm96/dku-aegis-library/model"
	"github.com/ilcm96/dku-aegis-library/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	SignUp(user *model.User) error
	SignIn(user *model.User) (token string, err error)
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
	if err == nil { // 가입된 아이디가 있다면 err == nil 이다 (없다면 *ErrNotFound 가 발생한다)
		return errors.New("ERR_ALREADY_EXISTS")
	}

	// 비밀번호 암호화
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return us.userRepo.Create(user)
}

func (us *userService) SignIn(user *model.User) (token string, err error) {
	queriedUser, err := us.userRepo.FindUserById(user.Id)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(queriedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	return makeJwt(queriedUser)
}

func makeJwt(user *ent.User) (string, error) {
	//TODO Jwt 유효시간 환경변수에서 가져오기
	exp := time.Now().Add(time.Hour * 24).Unix()
	claims := jwt.MapClaims{
		"id":   user.ID,
		"name": user.Name,
		"exp":  exp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//TODO Jwt 서명 시그니쳐 환경변수에서 가져오기
	t, err := token.SignedString([]byte("jwt-secret"))
	if err != nil {
		return "", err
	}

	return t, nil
}
