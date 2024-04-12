package handler

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ilcm96/dku-aegis-library/db"
	"github.com/ilcm96/dku-aegis-library/ent/user"
	"github.com/ilcm96/dku-aegis-library/model"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *fiber.Ctx) error {
	u := new(model.User)
	if err := c.BodyParser(u); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString("ERR_INVALID_REQUEST")
	}

	// Validate inputs
	if !((32000000 <= u.Id) && (u.Id <= 32999999)) {
		return c.Status(fiber.StatusBadRequest).SendString("ERR_INVALID_ID")
	} else if !(8 <= len(u.Password) && len(u.Password) < 20) {
		return c.Status(fiber.StatusBadRequest).SendString("ERR_INVALID_PASSWORD")
	} else if !(2 <= len(u.Name)) {
		return c.Status(fiber.StatusBadRequest).SendString("ERR_INVALID_NAME")
	}

	// Check whether ID is unique or not
	count, err := db.Client.User.Query().
		Where(user.ID(u.Id)).
		Count(context.Background())
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString("ERR_INVALID_REQUEST")
	}
	if count != 0 {
		return c.Status(fiber.StatusBadRequest).SendString("ERR_ID_EXIST")
	}

	// Hash password using bcrypt
	hashedPw, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString("ERR_INVALID_REQUEST")
	}

	// Run query
	newUser, err := db.Client.User.Create().
		SetID(u.Id).
		SetPassword(string(hashedPw)).
		SetName(u.Name).
		Save(context.Background())
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString("ERR_INVALID_REQUEST")
	}

	return c.Status(fiber.StatusCreated).SendString(strconv.Itoa(newUser.ID))
}

func Login(c *fiber.Ctx) error {
	u := new(model.User)
	if err := c.BodyParser(u); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString("ERR_INVALID_REQUEST")
	}

	loginUser, err := db.Client.User.Query().
		Where(user.ID(u.Id)).
		First(context.Background())
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString("ERR_INVALID_REQUEST")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(loginUser.Password), []byte(u.Password)); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString("ERR_INVALID_REQUEST")
	}

	// JWT
	exp := time.Now().Add(time.Hour * 1440)
	claims := jwt.MapClaims{
		"id":   loginUser.ID,
		"name": loginUser.Name,
		"exp":  exp.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("jwt-secret"))
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString("ERR_INVALID_REQUEST")
	}

	// Cookie
	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = t
	cookie.Expires = exp
	cookie.HTTPOnly = true
	c.Cookie(cookie)

	return c.SendStatus(fiber.StatusOK)
}
