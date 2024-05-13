package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ilcm96/dku-aegis-library/repository"
	"github.com/ilcm96/dku-aegis-library/service"
	"github.com/ilcm96/dku-aegis-library/util"
	"log/slog"
)

type BookController struct {
	bookService   service.BookService
	logRepository repository.LogRepository
}

func NewBookController(bookService service.BookService, logRepository repository.LogRepository) *BookController {
	return &BookController{
		bookService:   bookService,
		logRepository: logRepository,
	}
}

func (bc *BookController) BorrowBook(c *fiber.Ctx) error {
	userId := c.Context().UserValue("user-id").(int)
	requestId := c.Context().UserValue("request-id").(string)

	borrowReq := new(bookReq)
	if err := c.BodyParser(borrowReq); err != nil {
		util.LogErrWithReqId(c, err)
		return c.SendStatus(fiber.StatusBadRequest)
	}
	bookId := borrowReq.BookId

	book, err := bc.bookService.BorrowBook(bookId, userId)
	if err != nil {
		if err.Error() == "USER_ALREADY_BORROW" {
			util.LogErrWithReqId(c, err)
			return c.SendStatus(fiber.StatusForbidden)
		} else {
			util.LogErrWithReqId(c, err)
			return c.SendStatus(fiber.StatusBadRequest)
		}
	}

	slog.Info("borrow book", "user-id", userId, "request-id", requestId, "book-id", bookId)
	_ = bc.logRepository.LogBook("BORROW", userId, bookId, book.Title, requestId)
	return c.SendStatus(fiber.StatusCreated)
}

func (bc *BookController) ReturnBook(c *fiber.Ctx) error {
	userId := c.Context().UserValue("user-id").(int)
	requestId := c.Context().UserValue("request-id").(string)

	returnReq := new(bookReq)
	if err := c.BodyParser(returnReq); err != nil {
		util.LogErrWithReqId(c, err)
		return c.SendStatus(fiber.StatusBadRequest)
	}
	bookId := returnReq.BookId

	book, err := bc.bookService.ReturnBook(bookId, userId)
	if err != nil {
		util.LogErrWithReqId(c, err)
		return c.SendStatus(fiber.StatusBadRequest)
	}
	slog.Info("return book", "user-id", userId, "request-id", requestId, "book-id", bookId)
	_ = bc.logRepository.LogBook("RETURN", userId, bookId, book.Title, requestId)
	return c.SendStatus(fiber.StatusCreated)
}

type bookReq struct {
	BookId int `json:"bookId"`
}
