package mongorepo

import (
	"github.com/ahmadaidin/echoscratch/domain/entity"
	"github.com/ahmadaidin/echoscratch/domain/repository"
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

func (r *bookRepository) FindAll() (books []entity.Book, err error) {
	return
}
