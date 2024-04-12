package repository

import (
	"context"
	"github.com/ilcm96/dku-aegis-library/ent"
)

type BookRepository interface {
	FindAllBook() ([]*ent.Book, error)
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
