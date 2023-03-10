package mongorepo

import (
	"context"

	"github.com/ahmadaidin/gonostic/domain/entity"
	"github.com/ahmadaidin/gonostic/domain/model/book"
	"github.com/ahmadaidin/gonostic/domain/repository"
	"github.com/ahmadaidin/gonostic/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

type bookRepository struct {
	store *mongo.Database
}

func NewBookRepository(store *mongo.Database) repository.BookRepository {
	return &bookRepository{
		store: store,
	}
}

func createError() error {
	return errors.WithCallerInfo(errors.New("1st err"))
}

func (r *bookRepository) FindAll(ctx context.Context, opt ...book.FindOptions) (books []entity.Book, err error) {
	books = []entity.Book{
		{
			ID:    "1",
			Title: "The Lord of the Rings",
			Summary: `The Lord of the Rings is an epic high fantasy novel written by English author and scholar J. R. R. Tolkien.
The story began as a sequel to Tolkien's 1937 fantasy novel The Hobbit, but eventually developed into a much larger work.
Written in stages between 1937 and 1949, The Lord of the Rings is one of the best-selling novels ever written, with over 150 million copies sold.`,
		},
	}
	err = createError()
	return books, errors.WrapWithError(errors.WithCallerInfo(errors.New("2nd err")), err)
}
