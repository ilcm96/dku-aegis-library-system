package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ilcm96/dku-aegis-library/ent"
	user2 "github.com/ilcm96/dku-aegis-library/ent/user"
	"time"

	"github.com/ilcm96/dku-aegis-library/model"
	"github.com/ilcm96/dku-aegis-library/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	SignUp(user *model.User) error
	SignIn(user *model.User) (token string, err error)
	Withdraw(userId int) error
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
	u, err := us.userRepo.FindUserById(user.Id)
	if err != nil {
		if ent.IsNotFound(err) { // 에러가 NotFoundError 인 경우 신규회원이다
			if err := user.HashPassword(); err != nil {
				return err
			}
			return us.userRepo.Create(user)
		}
		return err // NotFoundError 가 아닌 경우 DB 자체 에러이다
	}

	if u.Status == user2.StatusWITHDRAW { // 탈퇴 회원인 경우
		// 비밀번호 암호화
		if err := user.HashPassword(); err != nil {
			return err
		}
		return us.userRepo.Update(user)
	} else { // 탈퇴 회원이 아닌 경우
		return errors.New("ERR_ALREADY_EXISTS")
	}
}

func (us *userService) SignIn(user *model.User) (token string, err error) {
	queriedUser, err := us.userRepo.FindUserById(user.Id)
	if err != nil {
		return "", err
	}

	if queriedUser.Status == user2.StatusPENDING {
		return "", errors.New("ERR_PENDING_USER")
	}

	err = bcrypt.CompareHashAndPassword([]byte(queriedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	return makeJwt(queriedUser)
}

func (us *userService) Withdraw(userId int) error {
	return us.userRepo.Withdraw(userId)
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
