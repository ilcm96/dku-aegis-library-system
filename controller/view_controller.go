package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ilcm96/dku-aegis-library/ent"
	"github.com/ilcm96/dku-aegis-library/repository"
	"github.com/ilcm96/dku-aegis-library/util"
	"os"
	"strconv"
)

type ViewController struct {
	userRepository    repository.UserRepository
	bookRepository    repository.BookRepository
	logRepository     repository.LogRepository
	bookReqRepository repository.BookReqRepository
}

func NewViewController(
	userRepository repository.UserRepository,
	bookRepository repository.BookRepository,
	logRepository repository.LogRepository,
	bookReqRepository repository.BookReqRepository,
) *ViewController {
	return &ViewController{
		userRepository:    userRepository,
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
	return c.Render("app/index", fiber.Map{
		"BookList": bookList,
		"IsAdmin":  c.Context().UserValue("is-admin").(bool),
		"IsProd":   os.Getenv("PRODUCTION") == "true",
	})
}

func (vc *ViewController) BookDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return c.Render("app/404", nil)
	}

	b, err := vc.bookRepository.FindBookById(intId)
	if err != nil {
		return c.Render("app/404", nil)
	}

	bookLogs, err := vc.logRepository.FilterByBookId(intId)
	return c.Render("book/book_detail", fiber.Map{
		"Book":    b,
		"BookLog": formatBookLog(bookLogs),
		"IsProd":  os.Getenv("PRODUCTION") == "true",
	})
}

func (vc *ViewController) MyPage(c *fiber.Ctx) error {
	userId := c.Context().UserValue("user-id").(int)
	bookLogs, _ := vc.logRepository.FilterByUserId(userId)
	borrowBooks, _ := vc.bookRepository.FindBooksByUserId(userId)

	return c.Render("app/mypage", fiber.Map{
		"BorrowBooks": borrowBooks,
		"BookLog":     formatBookLog(bookLogs),
		"IsAdmin":     c.Context().UserValue("is-admin").(bool),
	})
}

// -----------------
// --- Auth View ---
// -----------------

func (vc *ViewController) SignUp(c *fiber.Ctx) error {
	return c.Render("sign/signup", fiber.Map{})
}

func (vc *ViewController) Signin(c *fiber.Ctx) error {
	return c.Render("sign/signin", fiber.Map{})
}

// -------------------
// --- Search View ---
// -------------------

func (vc *ViewController) Search(c *fiber.Ctx) error {
	return c.Render("search/search", fiber.Map{})
}

func (vc *ViewController) SearchResult(c *fiber.Ctx) error {
	var req searchReq
	if err := c.BodyParser(&req); err != nil {
		util.LogErrWithReqId(c, err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if req.Search == "" {
		bookList, _ := vc.bookRepository.FindAllBook()
		return c.Render("search/search_result", fiber.Map{
			"BookList": bookList,
		}, "")
	}
	bookList, err := vc.bookRepository.SearchBook(req.Search)
	if err != nil {
		util.LogErrWithReqId(c, err)
	}

	return c.Render("search/search_result", fiber.Map{
		"BookList": bookList,
	}, "")
}

// -------------------
// --- Book Reqeust ---
// -------------------

func (vc *ViewController) BookRequest(c *fiber.Ctx) error {
	return c.Render("book/book_request", nil)
}

func (vc *ViewController) BookRequestHistory(c *fiber.Ctx) error {
	userId := c.Context().UserValue("user-id").(int)
	bookReqHistory, err := vc.bookReqRepository.FindBookReqByUserId(userId)
	if err != nil {
		util.LogErrWithReqId(c, err)
	}
	return c.Render("book/book_req_history", fiber.Map{
		"BookReqHistory": bookReqHistory,
	})
}

// -------------------
// --- Search View ---
// -------------------

func (vc *ViewController) Admin(c *fiber.Ctx) error {
	return c.Render("admin/admin", fiber.Map{})
}

func (vc *ViewController) AdminBook(c *fiber.Ctx) error {
	books, err := vc.bookRepository.FindAllBook()
	if err != nil {
		util.LogErrWithReqId(c, err)
	}

	logs, err := vc.logRepository.FindAllLogs()
	if err != nil {
		util.LogErrWithReqId(c, err)
	}

	return c.Render("admin/admin_book", fiber.Map{
		"Books": books,
		"Logs":  formatBookLog(logs),
	})
}

func (vc *ViewController) AdminBookCreate(c *fiber.Ctx) error {
	return c.Render("admin/admin_book_create", fiber.Map{})
}

func (vc *ViewController) AdminBookDetail(c *fiber.Ctx) error {
	bookId := c.Params("id")
	bookIdInt, err := strconv.Atoi(bookId)
	if err != nil {
		return c.Render("app/404", nil)
	}

	book, err := vc.bookRepository.FindBookById(bookIdInt)
	if err != nil {
		util.LogErrWithReqId(c, err)
	}

	return c.Render("admin/admin_book_detail", fiber.Map{
		"Book":   book,
		"IsProd": os.Getenv("PRODUCTION") == "true",
	})
}

func (vc *ViewController) AdminUser(c *fiber.Ctx) error {
	users, err := vc.userRepository.FindAllUser()
	if err != nil {
		util.LogErrWithReqId(c, err)
	}
	return c.Render("admin/admin_user", fiber.Map{
		"Users": users,
	})
}

func (vc *ViewController) AdminUserDetail(c *fiber.Ctx) error {
	userId := c.Params("id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		util.LogErrWithReqId(c, err)
		return c.Render("app/404", nil)
	}
	user, err := vc.userRepository.FindUserById(userIdInt)
	if err != nil {
		util.LogErrWithReqId(c, err)
	}

	return c.Render("admin/admin_user_detail", user)
}

func (vc *ViewController) AdminRequest(c *fiber.Ctx) error {
	bookReqs, err := vc.bookReqRepository.FindAllBookReq()
	if err != nil {
		util.LogErrWithReqId(c, err)
	}
	return c.Render("admin/admin_book_req", fiber.Map{
		"BookReqs": bookReqs,
	})
}

func (vc *ViewController) AdminRequestDetail(c *fiber.Ctx) error {
	bookReqId := c.Params("id")
	bookReqIdInt, err := strconv.Atoi(bookReqId)
	if err != nil {
		util.LogErrWithReqId(c, err)
		return c.Render("app/404", nil)
	}
	req, err := vc.bookReqRepository.FindBookReqById(bookReqIdInt)
	if err != nil {
		util.LogErrWithReqId(c, err)
	}

	return c.Render("admin/admin_book_req_detail", req)
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
