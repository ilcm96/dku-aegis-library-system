package service

import (
	"context"
	"errors"
	"github.com/h2non/filetype"
	"github.com/ilcm96/dku-aegis-library/ent"
	"github.com/ilcm96/dku-aegis-library/model"
	"github.com/ilcm96/dku-aegis-library/repository"
	"io"
	"mime/multipart"
)

type BookService interface {
	BorrowBook(bookId int, userId int) (*ent.Book, error)
	ReturnBook(bookId int, userId int) (*ent.Book, error)
	CreateBook(book *model.Book) (bookId int, err error)
	UpdateBook(bookId int, book *model.Book) error
	UpdateBookCover(bookId int, file multipart.File, filesize int64) error
	DeleteBook(bookId int) error
}

type bookService struct {
	bookRepo repository.BookRepository
}

func NewBookService(bookRepo repository.BookRepository) BookService {
	return &bookService{
		bookRepo: bookRepo,
	}
}

func (bs *bookService) BorrowBook(bookId int, userId int) (*ent.Book, error) {
	book, err := bs.bookRepo.FindBookById(bookId)
	if err != nil {
		return nil, err
	}
	if book.Borrow >= book.Quantity {
		return nil, errors.New("BORROW_EXCEED_QUANTITY")
	}

	users, err := book.QueryUser().All(context.Background())
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		if user.ID == userId {
			return nil, errors.New("USER_ALREADY_BORROW")
		}
	}

	return bs.bookRepo.BorrowBook(bookId, userId)
}

func (bs *bookService) ReturnBook(bookId int, userId int) (*ent.Book, error) {
	book, err := bs.bookRepo.FindBookById(bookId)
	if err != nil {
		return nil, err
	}

	users, err := book.QueryUser().All(context.Background())
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		if user.ID == userId {
			return bs.bookRepo.ReturnBook(bookId, userId)
		}
	}

	return nil, errors.New("USER_DID_NOT_BORROW_BOOK")
}

func (bs *bookService) CreateBook(book *model.Book) (bookId int, err error) {
	b, err := bs.bookRepo.CreateBook(book)
	if err != nil {
		return -1, err
	}
	return b.ID, err
}

func (bs *bookService) UpdateBook(bookId int, book *model.Book) error {
	return bs.bookRepo.UpdateBook(bookId, book)
}

func (bs *bookService) UpdateBookCover(bookId int, file multipart.File, filesize int64) error {
	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	fileType, err := filetype.Get(bytes)
	if err != nil {
		return err
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	if filetype.IsImage(bytes) {
		return bs.bookRepo.UpdateBookCover(bookId, file, fileType.Extension, filesize, fileType.MIME.Value)
	} else {
		return errors.New("ERR_NOT_IMAGE")
	}
}

func (bs *bookService) DeleteBook(bookId int) error {
	return bs.bookRepo.DeleteBook(bookId)
}
