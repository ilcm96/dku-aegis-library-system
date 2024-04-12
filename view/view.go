package view

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/ilcm96/dku-aegis-library/db"
)

func Index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		//TODO Redis 등으로 캐싱하기
		"BookList": db.Client.Book.Query().WithCategory().AllX(context.Background()),
	})
}

func Login(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{})
}

func SignUp(c *fiber.Ctx) error {
	return c.Render("signup", fiber.Map{})
}
