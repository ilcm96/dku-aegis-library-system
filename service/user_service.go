package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/ilcm96/dku-aegis-library/ent"
	user2 "github.com/ilcm96/dku-aegis-library/ent/user"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"

	"github.com/ilcm96/dku-aegis-library/model"
	"github.com/ilcm96/dku-aegis-library/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	SignUp(user *model.User) error
	SignIn(user *model.User) (string, error)
	SignOut(sessId string) error
	Withdraw(userId int) error
}

type userService struct {
	userRepo    repository.UserRepository
	redisClient *redis.Client
}

func NewUserService(userRepo repository.UserRepository, redisClient *redis.Client) UserService {
	return &userService{
		userRepo:    userRepo,
		redisClient: redisClient,
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

func (us *userService) SignIn(user *model.User) (string, error) {
	queriedUser, err := us.userRepo.FindUserById(user.Id)
	if err != nil {
		return "", err
	}

	if queriedUser.Status == user2.StatusPENDING {
		return "", errors.New("ERR_PENDING_USER")
	} else if queriedUser.Status == user2.StatusWITHDRAW {
		return "", errors.New("ERR_WITHDRAW_USER")
	}

	err = bcrypt.CompareHashAndPassword([]byte(queriedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	sessId := uuid.New().String()

	if err = us.redisClient.Set(context.Background(), sessId, user.Id, 10*time.Minute).Err(); err != nil {
		return "", err
	}

	if err = us.redisClient.RPush(context.Background(), strconv.Itoa(user.Id), sessId).Err(); err != nil {
		return "", err
	}

	return sessId, nil
}

func (us *userService) SignOut(sessId string) error {
	userId, err := us.redisClient.Get(context.Background(), sessId).Result()
	if err != nil {
		return err
	}

	if err = us.redisClient.Del(context.Background(), sessId).Err(); err != nil {
		return err
	}

	if err = us.redisClient.LRem(context.Background(), userId, 0, sessId).Err(); err != nil {
		return err
	}

	return nil
}

func (us *userService) Withdraw(userId int) error {
	if err := us.redisClient.Del(context.Background(), strconv.Itoa(userId)).Err(); err != nil {
		return err
	}

	if err := us.userRepo.Withdraw(userId); err != nil {
		return err
	}

	return nil
}
