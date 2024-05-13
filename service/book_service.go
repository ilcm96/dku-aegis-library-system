package service

import (
	"context"
	"errors"
	"github.com/ilcm96/dku-aegis-library/repository"
)

type BookService interface {
	BorrowBook(bookId int, userId int) error
	ReturnBook(bookId int, userId int) error
}

type bookService struct {
	bookRepo repository.BookRepository
}

func NewBookService(bookRepo repository.BookRepository) BookService {
	return &bookService{
		bookRepo: bookRepo,
	}
}

func (bs *bookService) BorrowBook(bookId int, userId int) error {
	book, err := bs.bookRepo.FindBookById(bookId)
	if err != nil {
		return err
	}
	if book.Borrow >= book.Quantity {
		return errors.New("BORROW_EXCEED_QUANTITY")
	}

	users, err := book.QueryUser().All(context.Background())
	if err != nil {
		return err
	}
	for _, user := range users {
		if user.ID == userId {
			return errors.New("USER_ALREADY_BORROW")
		}
	}

	return bs.bookRepo.BorrowBook(bookId, userId)
}

func (bs *bookService) ReturnBook(bookId int, userId int) error {
	book, err := bs.bookRepo.FindBookById(bookId)
	if err != nil {
		return err
	}

	users, err := book.QueryUser().All(context.Background())
	if err != nil {
		return err
	}
	for _, user := range users {
		if user.ID == userId {
			return bs.bookRepo.ReturnBook(bookId, userId)
		}
	}

	return errors.New("USER_DID_NOT_BORROW_BOOK")
}
