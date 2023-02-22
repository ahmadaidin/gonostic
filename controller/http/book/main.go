package book

import (
	"net/http"

	"github.com/ahmadaidin/echoscratch/domain/model"
	"github.com/ahmadaidin/echoscratch/domain/repository"
	"github.com/ahmadaidin/echoscratch/pkg/errcode"
	"github.com/labstack/echo/v4"
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

func (ctr *BookController) FindAll(c echo.Context) (err error) {
	filter := &model.BookFilter{}
	if err = c.Bind(filter); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errcode.ErrBadParam.Error())
	}
	books, err := ctr.bookRepo.FindAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, books)
}
