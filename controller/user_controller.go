package controller

import (
	"errors"
	"github.com/ilcm96/dku-aegis-library/ent"
	"github.com/ilcm96/dku-aegis-library/ent/user"
	"github.com/ilcm96/dku-aegis-library/util"
	"golang.org/x/crypto/bcrypt"
	"strconv"
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

	sessId, err := uc.userService.SignIn(user)
	if err != nil {
		util.LogErrWithReqId(c, err)
		if ent.IsNotFound(err) || errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) || err.Error() == "ERR_WITHDRAW_USER" {
			return c.SendStatus(fiber.StatusUnauthorized)
		} else if err.Error() == "ERR_PENDING_USER" {
			return c.SendStatus(fiber.StatusForbidden)
		} else {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	}

	c.Cookie(&fiber.Cookie{
		Name:     "session_id",
		Value:    sessId,
		Path:     "/",
		Expires:  time.Now().Add(10 * time.Minute),
		HTTPOnly: true,
		SameSite: "Lax",
	})

	return c.SendStatus(fiber.StatusOK)
}

func (uc *UserController) SignOut(c *fiber.Ctx) error {
	if err := uc.userService.SignOut(c.Cookies("session_id")); err != nil {
		util.LogErrWithReqId(c, err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	removeSessionCookie(c)

	return c.SendStatus(fiber.StatusOK)
}

func (uc *UserController) Withdraw(c *fiber.Ctx) error {
	userId := c.Context().UserValue("user-id").(int)

	if err := uc.userService.Withdraw(userId); err != nil {
		util.LogErrWithReqId(c, err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	removeSessionCookie(c)

	return c.SendStatus(fiber.StatusOK)
}

func (uc *UserController) ChangeStatus(c *fiber.Ctx) error {
	isAdmin := c.Context().UserValue("is-admin").(bool)
	currentSessionUserId := c.Context().UserValue("user-id").(int)

	userId := c.Params("id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		util.LogErrWithReqId(c, err)
		return c.SendStatus(fiber.StatusNotFound)
	}

	if isAdmin {
		if currentSessionUserId == userIdInt {
			return c.SendStatus(fiber.StatusForbidden)
		}
	}

	var status struct {
		Status user.Status `json:"status"`
	}

	if err = c.BodyParser(&status); err != nil {
		util.LogErrWithReqId(c, err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err = uc.userService.ChangeStatus(userIdInt, status.Status); err != nil {
		if err.Error() == "ERR_WITHDRAW_USER" {
			util.LogErrWithReqId(c, err)
			return c.SendStatus(fiber.StatusBadRequest)
		}
		util.LogErrWithReqId(c, err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusCreated)
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

func removeSessionCookie(c *fiber.Ctx) {
	c.Cookie(&fiber.Cookie{
		Name:    "session_id",
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
	})
}
