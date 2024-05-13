package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ilcm96/dku-aegis-library/service"
	"github.com/ilcm96/dku-aegis-library/util"
)

type BookController struct {
	bookService service.BookService
}

func NewBookController(bookService service.BookService) *BookController {
	return &BookController{
		bookService: bookService,
	}
}

func (bc *BookController) BorrowBook(c *fiber.Ctx) error {
	borrowReq := new(bookReq)
	if err := c.BodyParser(borrowReq); err != nil {
		util.LogErrWithReqId(c, err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err := bc.bookService.BorrowBook(borrowReq.BookId, c.Context().UserValue("user-id").(int))
	if err != nil {
		if err.Error() == "USER_ALREADY_BORROW" {
			util.LogErrWithReqId(c, err)
			return c.SendStatus(fiber.StatusForbidden)
		} else {
			util.LogErrWithReqId(c, err)
			return c.SendStatus(fiber.StatusBadRequest)
		}
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (bc *BookController) ReturnBook(c *fiber.Ctx) error {
	returnReq := new(bookReq)
	if err := c.BodyParser(returnReq); err != nil {
		util.LogErrWithReqId(c, err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err := bc.bookService.ReturnBook(returnReq.BookId, c.Context().UserValue("user-id").(int))
	if err != nil {
		util.LogErrWithReqId(c, err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusCreated)
}

type bookReq struct {
	BookId int `json:"bookId"`
}
