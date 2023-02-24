package book

import (
	"net/http"

	"github.com/ahmadaidin/echoscratch/core"
	"github.com/ahmadaidin/echoscratch/core/adapter"
	"github.com/ahmadaidin/echoscratch/domain/model"
	"github.com/ahmadaidin/echoscratch/domain/repository"
	"github.com/ahmadaidin/echoscratch/pkg"
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
	if err = c.Bind(filter); err != nil {
		return adapter.NewHTTPError(http.StatusBadRequest, pkg.NewError(pkg.ErrBadParam, err, pkg.MsgErrBadParam))
	}
	books, err := ctr.bookRepo.FindAll()
	if err != nil {
		return adapter.NewHTTPError(http.StatusInternalServerError, pkg.NewError(pkg.ErrUnexpected, err))
	}
	return c.JSON(http.StatusOK, books)
}
