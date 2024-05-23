package repository

import (
	"context"
	"github.com/ilcm96/dku-aegis-library/ent"
	"github.com/ilcm96/dku-aegis-library/ent/bookrequest"
	"github.com/ilcm96/dku-aegis-library/model"
)

type BookReqRepository interface {
	FindAllBookReq() ([]*ent.BookRequest, error)
	FindBookReqByUserId(userId int) ([]*ent.BookRequest, error)
	CreateBookReq(bookReq *model.BookReq) error
	DeleteBookReq(id int) error
}

type bookReqRepository struct {
	client *ent.Client
}

func NewBookReqRepository(client *ent.Client) BookReqRepository {
	return &bookReqRepository{
		client: client,
	}
}

func (br *bookReqRepository) FindAllBookReq() ([]*ent.BookRequest, error) {
	return br.client.BookRequest.Query().
		All(context.Background())
}

func (br *bookReqRepository) FindBookReqByUserId(userId int) ([]*ent.BookRequest, error) {
	return br.client.BookRequest.Query().
		Where(bookrequest.UserID(userId)).
		All(context.Background())
}

func (br *bookReqRepository) CreateBookReq(bookReq *model.BookReq) error {
	_, err := br.client.BookRequest.Create().
		SetUserID(bookReq.UserId).
		SetTitle(bookReq.Title).
		SetAuthor(bookReq.Author).
		SetPublisher(bookReq.Publisher).
		SetReason(bookReq.Reason).
		Save(context.Background())

	return err
}

func (br *bookReqRepository) DeleteBookReq(id int) error {
	return br.client.BookRequest.DeleteOneID(id).
		Exec(context.Background())
}
