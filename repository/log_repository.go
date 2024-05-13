package repository

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/ilcm96/dku-aegis-library/ent"
	"github.com/ilcm96/dku-aegis-library/ent/booklog"
)

type LogRepository interface {
	LogBook(action string, userId int, bookId int, bookTitle string, requestId string) error
	FilterByUserId(userId int) ([]*ent.BookLog, error)
}

type logRepository struct {
	client *ent.Client
}

func NewLogRepository(client *ent.Client) LogRepository {
	return &logRepository{
		client: client,
	}
}

func (lr *logRepository) LogBook(action string, userId int, bookId int, bookTitle string, requestId string) error {
	_, err := lr.client.BookLog.Create().
		SetAction(booklog.Action(action)).
		SetUserID(userId).
		SetBookID(bookId).
		SetBookTitle(bookTitle).
		SetRequestID(requestId).
		Save(context.Background())

	return err
}

func (lr *logRepository) FilterByUserId(userId int) ([]*ent.BookLog, error) {
	return lr.client.BookLog.Query().
		Where(booklog.UserID(userId)).
		Order(booklog.ByID(sql.OrderDesc())).
		All(context.Background())
}
