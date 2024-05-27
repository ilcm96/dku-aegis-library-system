package controller

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/ilcm96/dku-aegis-library/ent/booklog"
	"github.com/ilcm96/dku-aegis-library/model"
	"github.com/ilcm96/dku-aegis-library/repository"
	"github.com/ilcm96/dku-aegis-library/service"
	"github.com/ilcm96/dku-aegis-library/util"
	"log/slog"
	"strconv"
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
	_ = bc.logRepository.LogBook(booklog.ActionBORROW, userId, bookId, book.Title, requestId)
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
	_ = bc.logRepository.LogBook(booklog.ActionRETURN, userId, bookId, book.Title, requestId)
	return c.SendStatus(fiber.StatusCreated)
}

func (bc *BookController) AdminCreateBook(c *fiber.Ctx) error {
	b := new(model.Book)
	if err := c.BodyParser(b); err != nil {
		util.LogErrWithReqId(c, err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if !b.Validate() {
		util.LogErrWithReqId(c, errors.New("invalid book"))
		return c.SendStatus(fiber.StatusBadRequest)
	}

	bookId, err := bc.bookService.CreateBook(b)
	if err != nil {
		util.LogErrWithReqId(c, err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.Status(fiber.StatusCreated).SendString(strconv.Itoa(bookId))
}

func (bc *BookController) AdminUpdateBook(c *fiber.Ctx) error {
	bookId := c.Params("id")
	bookIdInt, err := strconv.Atoi(bookId)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	b := new(model.Book)

	if err = c.BodyParser(b); err != nil {
		util.LogErrWithReqId(c, err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err = bc.bookService.UpdateBook(bookIdInt, b); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (bc *BookController) AdminUpdateBookCover(c *fiber.Ctx) error {
	bookId := c.Params("id")
	bookIdInt, err := strconv.Atoi(bookId)
	if err != nil {
		util.LogErrWithReqId(c, err)
		return c.SendStatus(fiber.StatusNotFound)
	}

	file, err := c.FormFile("file")
	if err != nil {
		util.LogErrWithReqId(c, err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	f, err := file.Open()
	if err != nil {
		util.LogErrWithReqId(c, err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer f.Close()

	if err = bc.bookService.UpdateBookCover(bookIdInt, f, file.Size); err != nil {
		util.LogErrWithReqId(c, err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (bc *BookController) AdminDeleteBook(c *fiber.Ctx) error {
	bookId := c.Params("id")
	bookIdInt, err := strconv.Atoi(bookId)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	if err = bc.bookService.DeleteBook(bookIdInt); err != nil {
		util.LogErrWithReqId(c, err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

type bookReq struct {
	BookId int `json:"bookId"`
}
