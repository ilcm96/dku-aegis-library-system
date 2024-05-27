package repository

import (
	"context"
	"fmt"
	"github.com/ilcm96/dku-aegis-library/ent"
	"github.com/ilcm96/dku-aegis-library/ent/book"
	"github.com/ilcm96/dku-aegis-library/ent/user"
	"github.com/ilcm96/dku-aegis-library/model"
	"github.com/minio/minio-go/v7"
	"mime/multipart"
)

type BookRepository interface {
	FindAllBook() ([]*ent.Book, error)
	FindBookById(bookId int) (*ent.Book, error)
	FindBooksByUserId(userId int) ([]*ent.Book, error)
	BorrowBook(bookId int, userId int) (*ent.Book, error)
	ReturnBook(bookId int, userId int) (*ent.Book, error)
	SearchBook(query string) ([]*ent.Book, error)
	CreateBook(book *model.Book) (*ent.Book, error)
	UpdateBook(bookId int, book *model.Book) error
	UpdateBookCover(bookId int, file multipart.File, extension string, filesize int64, mime string) error
	DeleteBook(bookId int) error
}

type bookRepository struct {
	client *ent.Client
	minio  *minio.Client
}

func NewBookRepository(client *ent.Client, minio *minio.Client) BookRepository {
	return &bookRepository{
		client: client,
		minio:  minio,
	}
}

func (br *bookRepository) FindAllBook() ([]*ent.Book, error) {
	return br.client.Book.Query().
		All(context.Background())
}

func (br *bookRepository) FindBookById(bookId int) (*ent.Book, error) {
	return br.client.Book.Get(context.Background(), bookId)
}

func (br *bookRepository) FindBooksByUserId(userId int) ([]*ent.Book, error) {
	return br.client.Book.Query().
		Where(book.HasUserWith(user.ID(userId))).
		All(context.Background())
}

func (br *bookRepository) BorrowBook(bookId int, userId int) (*ent.Book, error) {
	return br.client.Book.UpdateOneID(bookId).
		AddBorrow(1).
		AddUserIDs(userId).
		Save(context.Background())
}

func (br *bookRepository) ReturnBook(bookId int, userId int) (*ent.Book, error) {
	return br.client.Book.UpdateOneID(bookId).
		AddBorrow(-1).
		RemoveUserIDs(userId).
		Save(context.Background())
}

func (br *bookRepository) SearchBook(query string) ([]*ent.Book, error) {
	return br.client.Book.Query().
		Where(book.Or(
			book.TitleContains(query),
			book.AuthorContains(query),
			book.PublisherContains(query),
			book.CategoryContains(query),
		)).
		All(context.Background())
}

func (br *bookRepository) CreateBook(book *model.Book) (*ent.Book, error) {
	return br.client.Book.Create().
		SetTitle(book.Title).
		SetAuthor(book.Author).
		SetPublisher(book.Publisher).
		SetCategory(book.Category).
		SetQuantity(book.Quantity).
		SetIsbn(book.Isbn).
		Save(context.Background())
}

func (br *bookRepository) UpdateBook(bookId int, book *model.Book) error {
	return br.client.Book.UpdateOneID(bookId).
		SetTitle(book.Title).
		SetAuthor(book.Author).
		SetPublisher(book.Publisher).
		SetCategory(book.Category).
		SetQuantity(book.Quantity).
		SetIsbn(book.Isbn).
		Exec(context.Background())
}

func (br *bookRepository) UpdateBookCover(bookId int, file multipart.File, extension string, filesize int64, mime string) error {
	filename := fmt.Sprintf("%d.%s", bookId, extension)
	_, err := br.minio.PutObject(
		context.Background(),
		"dku-aegis-library-system-cover-image",
		filename,
		file,
		filesize,
		minio.PutObjectOptions{ContentType: mime},
	)
	if err != nil {
		return err
	}

	return br.client.Book.Update().
		Where(book.ID(bookId)).
		SetCover(filename).
		Exec(context.Background())
}

func (br *bookRepository) DeleteBook(bookId int) error {
	return br.client.Book.DeleteOneID(bookId).
		Exec(context.Background())
}
