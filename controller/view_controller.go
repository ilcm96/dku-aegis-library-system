package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ilcm96/dku-aegis-library/repository"
)

type ViewController struct {
	bookRepository repository.BookRepository
}

func NewViewController(bookRepository repository.BookRepository) *ViewController {
	return &ViewController{
		bookRepository: bookRepository,
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
	bookList, _ := vc.bookRepository.FindBooksByUserId(c.Context().UserValue("user-id").(int))
	fmt.Println(len(bookList))
	return c.Render("mypage", fiber.Map{
		"BookList": bookList,
	})
}
