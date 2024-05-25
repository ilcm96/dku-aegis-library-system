package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ilcm96/dku-aegis-library/ent"
	"github.com/ilcm96/dku-aegis-library/repository"
	"github.com/ilcm96/dku-aegis-library/util"
	"strconv"
)

type ViewController struct {
	bookRepository    repository.BookRepository
	logRepository     repository.LogRepository
	bookReqRepository repository.BookReqRepository
}

func NewViewController(
	bookRepository repository.BookRepository,
	logRepository repository.LogRepository,
	bookReqRepository repository.BookReqRepository,
) *ViewController {
	return &ViewController{
		bookRepository:    bookRepository,
		logRepository:     logRepository,
		bookReqRepository: bookReqRepository,
	}
}

// ------------------------
// --- Application View ---
// ------------------------

func (vc *ViewController) Index(c *fiber.Ctx) error {
	bookList, _ := vc.bookRepository.FindAllBook()
	return c.Render("index", fiber.Map{
		"BookList": bookList,
	})
}

func (vc *ViewController) BookDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return c.Render("404", nil)
	}

	b, err := vc.bookRepository.FindBookById(intId)
	if err != nil {
		return c.Render("404", nil)
	}

	bookLogs, err := vc.logRepository.FilterByBookId(intId)
	return c.Render("book_detail", fiber.Map{
		"Book":    b,
		"BookLog": formatBookLog(bookLogs),
	})
}

func (vc *ViewController) MyPage(c *fiber.Ctx) error {
	userId := c.Context().UserValue("user-id").(int)
	bookLogs, _ := vc.logRepository.FilterByUserId(userId)
	borrowBooks, _ := vc.bookRepository.FindBooksByUserId(userId)

	return c.Render("mypage", fiber.Map{
		"BorrowBooks": borrowBooks,
		"BookLog":     formatBookLog(bookLogs),
	})
}

// -----------------
// --- Auth View ---
// -----------------

func (vc *ViewController) SignUp(c *fiber.Ctx) error {
	return c.Render("signup", fiber.Map{})
}

func (vc *ViewController) Signin(c *fiber.Ctx) error {
	return c.Render("signin", fiber.Map{})
}

// -------------------
// --- Search View ---
// -------------------

func (vc *ViewController) Search(c *fiber.Ctx) error {
	return c.Render("search", fiber.Map{})
}

func (vc *ViewController) SearchResult(c *fiber.Ctx) error {
	var req searchReq
	if err := c.BodyParser(&req); err != nil {
		util.LogErrWithReqId(c, err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if req.Search == "" {
		bookList, _ := vc.bookRepository.FindAllBook()
		return c.Render("search_result", fiber.Map{
			"BookList": bookList,
		}, "")
	}
	bookList, err := vc.bookRepository.SearchBook(req.Search)
	if err != nil {
		util.LogErrWithReqId(c, err)
	}

	return c.Render("search_result", fiber.Map{
		"BookList": bookList,
	}, "")
}

// -------------------
// --- Book Reqeust ---
// -------------------

func (vc *ViewController) BookRequest(c *fiber.Ctx) error {
	return c.Render("book_request", nil)
}

func (vc *ViewController) BookRequestHistory(c *fiber.Ctx) error {
	userId := c.Context().UserValue("user-id").(int)
	bookReqHistory, err := vc.bookReqRepository.FindBookReqByUserId(userId)
	if err != nil {
		util.LogErrWithReqId(c, err)
	}
	return c.Render("book_req_history", fiber.Map{
		"BookReqHistory": bookReqHistory,
	})
}

func formatBookLog(logs []*ent.BookLog) []bookLog {
	var bookLogs []bookLog
	for _, log := range logs {
		date := log.CreatedAt.Format("06-01-02 15:04")
		action := ""
		switch log.Action {
		case "BORROW":
			action = "대출"

		case "RETURN":
			action = "반납"
		}
		bookLogs = append(bookLogs, bookLog{log.BookID, log.UserID, log.BookTitle, date, action})
	}
	return bookLogs
}

type bookLog struct {
	BookId int
	UserId int
	Title  string
	Date   string
	Action string
}

type searchReq struct {
	Search string `json:"search"`
}
