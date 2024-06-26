package repository

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/ilcm96/dku-aegis-library/ent"
	"github.com/ilcm96/dku-aegis-library/ent/booklog"
)

type LogRepository interface {
	FindAllLogs() ([]*ent.BookLog, error)
	LogBook(action booklog.Action, userId int, bookId int, bookTitle string, requestId string) error
	FilterByUserId(userId int) ([]*ent.BookLog, error)
	FilterByBookId(bookId int) ([]*ent.BookLog, error)
}

type logRepository struct {
	client *ent.Client
}

func NewLogRepository(client *ent.Client) LogRepository {
	return &logRepository{
		client: client,
	}
}

func (lr *logRepository) FindAllLogs() ([]*ent.BookLog, error) {
	return lr.client.BookLog.Query().
		All(context.Background())
}

func (lr *logRepository) LogBook(action booklog.Action, userId int, bookId int, bookTitle string, requestId string) error {
	_, err := lr.client.BookLog.Create().
		SetAction(action).
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

func (lr *logRepository) FilterByBookId(bookId int) ([]*ent.BookLog, error) {
	return lr.client.BookLog.Query().
		Where(booklog.BookID(bookId)).
		All(context.Background())
}
