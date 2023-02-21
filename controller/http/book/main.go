package book

import (
	"net/http"

	"github.com/ahmadaidin/echoscratch/domain/model"
	"github.com/ahmadaidin/echoscratch/domain/repository"
	"github.com/ahmadaidin/echoscratch/pkg/errcode"
	"github.com/labstack/echo/v4"
)

type bookController struct {
	bookRepo repository.BookRepository
}

func NewBookController(
	bookRepo repository.BookRepository,
) *bookController {
	return &bookController{
		bookRepo: bookRepo,
	}
}

func (ctr *bookController) FindAll(c echo.Context) (err error) {
	filter := &model.BookFilter{}
	if err = c.Bind(filter); err != nil {
		c.Echo().Logger.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, errcode.ErrBadParam.Error())
	}
	_, err = ctr.bookRepo.FindAll()
	return
}
