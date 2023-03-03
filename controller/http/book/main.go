package book

import (
	"net/http"

	"github.com/ahmadaidin/gonostic/core"
	"github.com/ahmadaidin/gonostic/domain/model"
	"github.com/ahmadaidin/gonostic/domain/repository"
	"github.com/ahmadaidin/gonostic/pkg"
)

type BookController struct {
	bookRepo repository.BookRepository
}

func NewBookController(
	bookRepo repository.BookRepository,
) *BookController {
	return &BookController{
		bookRepo: bookRepo,
	}
}

// @Summary Find all books
// @Description Find all books
// @Tags Book
// @Accept  json
// @Produce  json
// @Success 200 {object} []entity.Book
// @Router /books [get]
func (ctr *BookController) FindAll(c core.Context) (err error) {
	filter := &model.BookFilter{}
	if err = c.QueryParser(filter); err != nil {
		return c.HttpError(http.StatusBadRequest, pkg.NewError(pkg.ErrBadParam, err, pkg.MsgErrBadParam))
	}
	books, err := ctr.bookRepo.FindAll()
	if err != nil {
		return c.HttpError(http.StatusInternalServerError, pkg.NewError(pkg.ErrUnexpected, err))
	}
	return c.SendJson(http.StatusOK, books)
}
