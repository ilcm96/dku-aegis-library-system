package controller

import (
	"errors"
	"github.com/ilcm96/dku-aegis-library/ent"
	"github.com/ilcm96/dku-aegis-library/util"
	"golang.org/x/crypto/bcrypt"
	"time"
	"unicode/utf8"

	"github.com/gofiber/fiber/v2"
	"github.com/ilcm96/dku-aegis-library/model"
	"github.com/ilcm96/dku-aegis-library/service"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) SignUp(c *fiber.Ctx) error {
	user := new(model.User)
	if err := c.BodyParser(&user); err != nil {
		util.LogErrWithReqId(c, err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// 요청 검증
	if err := validate(user); err != nil {
		util.LogErrWithReqId(c, err)
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// 회원가입
	err := uc.userService.SignUp(user)
	if err != nil {
		util.LogErrWithReqId(c, err)
		if err.Error() == "ERR_ALREADY_EXISTS" {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (uc *UserController) SignIn(c *fiber.Ctx) error {
	user := new(model.User)
	if err := c.BodyParser(&user); err != nil {
		util.LogErrWithReqId(c, err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// JWT 생성
	token, err := uc.userService.SignIn(user)
	if err != nil {
		util.LogErrWithReqId(c, err)
		if ent.IsNotFound(err) || errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return c.SendStatus(fiber.StatusUnauthorized)
		} else if err.Error() == "ERR_PENDING_USER" {
			return c.SendStatus(fiber.StatusForbidden)
		} else {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	}

	// 쿠키 설정
	cookie := generateCookie(token)
	//fmt.Printf("%+v\n", cookie)
	c.Cookie(cookie)
	return c.SendStatus(fiber.StatusOK)
}

func (uc *UserController) Withdraw(c *fiber.Ctx) error {
	userId := c.Context().UserValue("user-id").(int)
	err := uc.userService.Withdraw(userId)
	if err != nil {
		util.LogErrWithReqId(c, err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}

func validate(user *model.User) error {
	if !(user.Id >= 32000000 && user.Id <= 32999999) {
		return errors.New("ERR_INVALID_ID")
	} else if !(utf8.RuneCountInString(user.Password) >= 8 && utf8.RuneCountInString(user.Password) <= 20) {
		//TODO 영문인지 체크
		return errors.New("ERR_INVALID_PASSWORD")
	} else if !(utf8.RuneCountInString(user.Name) >= 2 && utf8.RuneCountInString(user.Name) <= 5) {
		//TODO 한글인지 체크
		return errors.New("ERR_INVALID_NAME")
	} else {
		return nil
	}
}

func generateCookie(token string) *fiber.Cookie {
	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(time.Hour * 24)
	cookie.HTTPOnly = true

	return cookie
}
