package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ilcm96/dku-aegis-library/model"
	"github.com/ilcm96/dku-aegis-library/repository"
	"github.com/ilcm96/dku-aegis-library/util"
)

type BookReqController struct {
	BookReqRepository repository.BookReqRepository
}

func NewBookReqController(bookReqRepository repository.BookReqRepository) *BookReqController {
	return &BookReqController{
		BookReqRepository: bookReqRepository,
	}
}

func (bc *BookReqController) CreateBookReq(c *fiber.Ctx) error {
	bookReq := new(model.BookReq)
	if err := c.BodyParser(bookReq); err != nil {
		util.LogErrWithReqId(c, err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	userId := c.Context().UserValue("user-id").(int)
	bookReq.UserId = userId

	if err := bc.BookReqRepository.CreateBookReq(bookReq); err != nil {
		util.LogErrWithReqId(c, err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusCreated)
}
