package book

import (
	"net/http"

	"github.com/ahmadaidin/echoscratch/domain/model"
	"github.com/labstack/echo/v4"
)

type bookController struct {
	echo *echo.Echo
}

func NewBookController(echo *echo.Echo) *bookController {
	return &bookController{echo}
}

func (ctr *bookController) FindAll(c echo.Context) (err error) {
	filter := &model.BookFilter{}
	if err = c.Bind(filter); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad filter", err.Error())
	}
	return
}

// Start start the services
func (ctr *bookController) Routes(path string) {
	r := ctr.echo.Group(path)
	r.GET("", ctr.FindAll)
}
