package repository

import (
	"context"
	"github.com/ilcm96/dku-aegis-library/ent"
)

type BookRepository interface {
	FindAllBook() ([]*ent.Book, error)
	FindBookById(bookId int) (*ent.Book, error)
	BorrowBook(bookId int, userId int) error
	ReturnBook(bookId int, userId int) error
}

type bookRepository struct {
	client *ent.Client
}

func NewBookRepository(client *ent.Client) BookRepository {
	return &bookRepository{
		client: client,
	}
}

func (br *bookRepository) FindAllBook() ([]*ent.Book, error) {
	return br.client.Book.Query().
		WithCategory().
		All(context.Background())
}

func (br *bookRepository) FindBookById(bookId int) (*ent.Book, error) {
	return br.client.Book.Get(context.Background(), bookId)
}

func (br *bookRepository) BorrowBook(bookId int, userId int) error {
	_, err := br.client.Book.UpdateOneID(bookId).
		AddBorrow(1).
		AddUserIDs(userId).
		Save(context.Background())
	return err
}

func (br *bookRepository) ReturnBook(bookId int, userId int) error {
	_, err := br.client.Book.UpdateOneID(bookId).
		AddBorrow(-1).
		RemoveUserIDs(userId).
		Save(context.Background())
	return err
}
