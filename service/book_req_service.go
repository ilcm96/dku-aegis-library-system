package service

import (
	"github.com/ilcm96/dku-aegis-library/repository"
)

type BookReqService interface {
}

type bookReqService struct {
	bookReqRepo repository.BookReqRepository
}
