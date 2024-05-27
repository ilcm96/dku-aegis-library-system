package repository

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/ilcm96/dku-aegis-library/ent"
	"github.com/ilcm96/dku-aegis-library/ent/bookrequest"
	"github.com/ilcm96/dku-aegis-library/model"
)

type BookReqRepository interface {
	FindAllBookReq() ([]*ent.BookRequest, error)
	FindBookReqByUserId(userId int) ([]*ent.BookRequest, error)
	FindBookReqById(id int) (*ent.BookRequest, error)
	CreateBookReq(bookReq *model.BookReq) error
	UpdateBookReqApproved(reqId int, approved bookrequest.Approved) error
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
		Order(bookrequest.ByID(sql.OrderDesc())).
		All(context.Background())
}

func (br *bookReqRepository) FindBookReqByUserId(userId int) ([]*ent.BookRequest, error) {
	return br.client.BookRequest.Query().
		Where(bookrequest.UserID(userId)).
		Order(bookrequest.ByID(sql.OrderDesc())).
		All(context.Background())
}

func (br *bookReqRepository) FindBookReqById(id int) (*ent.BookRequest, error) {
	return br.client.BookRequest.Get(context.Background(), id)
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

func (br *bookReqRepository) UpdateBookReqApproved(reqId int, approved bookrequest.Approved) error {
	return br.client.BookRequest.UpdateOneID(reqId).
		SetApproved(approved).
		Exec(context.Background())
}

func (br *bookReqRepository) DeleteBookReq(id int) error {
	return br.client.BookRequest.DeleteOneID(id).
		Exec(context.Background())
}
