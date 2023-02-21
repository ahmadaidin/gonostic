package repository

import "github.com/ahmadaidin/echoscratch/domain/entity"

type BookRepository interface {
	FindAll() (books []entity.Book, err error)
}
