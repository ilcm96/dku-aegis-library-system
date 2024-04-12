package controller

import (
	"log"
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
		log.Println("ERR: cannot parse user |", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// 요청 검증
	if !validate(user) {
		log.Println("ERR: validation failed")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// 회원가입
	err := uc.userService.SignUp(user)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusCreated)
}

func validate(user *model.User) bool {
	log.Println(user)
	if !(user.Id >= 32000000 && user.Id <= 32999999) {
		log.Println(1)
		return false
	} else if !(len(user.Password) >= 8 && len(user.Password) <= 20) {
		//TODO 영문인지 체크
		log.Println(2)
		return false
	} else if !(utf8.RuneCountInString(user.Name) >= 2 && utf8.RuneCountInString(user.Name) <= 5) {
		//TODO 한글인지 체크
		log.Println(3)
		return false
	} else {
		return true
	}
}
