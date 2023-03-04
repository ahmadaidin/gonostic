package repository

import (
	"context"

	"github.com/ahmadaidin/gonostic/domain/entity"
	"github.com/ahmadaidin/gonostic/domain/model/book"
)

type BookRepository interface {
	FindAll(ctx context.Context, opt ...book.FindOptions) (books []entity.Book, err error)
}
