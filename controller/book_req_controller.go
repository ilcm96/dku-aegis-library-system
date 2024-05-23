package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ilcm96/dku-aegis-library/ent"
	"github.com/ilcm96/dku-aegis-library/model"
	"github.com/ilcm96/dku-aegis-library/repository"
	"github.com/ilcm96/dku-aegis-library/util"
	"strconv"
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

func (bc *BookReqController) DeleteBookReq(c *fiber.Ctx) error {
	bookReqId := c.Params("id")
	bookReqIdInt, err := strconv.Atoi(bookReqId)
	if err != nil {
		util.LogErrWithReqId(c, err)
		return c.SendStatus(fiber.StatusNotFound)
	}

	if err := bc.BookReqRepository.DeleteBookReq(bookReqIdInt); err != nil {
		util.LogErrWithReqId(c, err)
		if ent.IsNotFound(err) {
			return c.SendStatus(fiber.StatusNotFound)
		} else {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	}

	return c.SendStatus(fiber.StatusOK)
}
