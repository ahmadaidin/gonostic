package repository

import "github.com/ahmadaidin/gonostic/domain/entity"

type BookRepository interface {
	FindAll() (books []entity.Book, err error)
}
