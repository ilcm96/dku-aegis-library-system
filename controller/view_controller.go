package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ilcm96/dku-aegis-library/repository"
)

type ViewController struct {
	bookRepository repository.BookRepository
	logRepository  repository.LogRepository
}

func NewViewController(bookRepository repository.BookRepository, logRepository repository.LogRepository) *ViewController {
	return &ViewController{
		bookRepository: bookRepository,
		logRepository:  logRepository,
	}
}

func (vc *ViewController) Index(c *fiber.Ctx) error {
	bookList, _ := vc.bookRepository.FindAllBook()
	return c.Render("index", fiber.Map{
		"BookList": bookList,
	})
}

func (vc *ViewController) SignUp(c *fiber.Ctx) error {
	return c.Render("signup", fiber.Map{})
}

func (vc *ViewController) Login(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{})
}

func (vc *ViewController) MyPage(c *fiber.Ctx) error {
	bookLog := []book{}
	userId := c.Context().UserValue("user-id").(int)
	logs, _ := vc.logRepository.FilterByUserId(userId)

	for _, log := range logs {
		date := log.CreatedAt.Format("2006-01-02 15:04")
		action := ""
		switch log.Action {
		case "BORROW":
			action = "대출"

		case "RETURN":
			action = "반납"
		}
		bookLog = append(bookLog, book{log.BookID, log.BookTitle, date, action})
	}

	borrowBooks, _ := vc.bookRepository.FindBooksByUserId(userId)

	return c.Render("mypage", fiber.Map{
		"BorrowBooks": borrowBooks,
		"BookLog":     bookLog,
	})
}

type book struct {
	BookId int
	Title  string
	Date   string
	Action string
}
