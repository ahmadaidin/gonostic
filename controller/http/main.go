package http

import (
	"context"
	"fmt"

	"github.com/go-playground/mold/v4"
	"github.com/go-playground/mold/v4/modifiers"
	"github.com/go-playground/validator/v10"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/ahmadaidin/echoscratch/config"
	"github.com/ahmadaidin/echoscratch/controller/http/book"
	"github.com/ahmadaidin/echoscratch/pkg/binder"
	"github.com/labstack/echo/v4" // we use echo version 4 here
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	_ "github.com/ahmadaidin/echoscratch/docs"
)

// customValidator to validate input request
type customValidator struct {
	validator *validator.Validate
	modifier  *mold.Transformer
	autoMod   bool
}

// validate using custom validator
func (cv customValidator) Validate(i interface{}) error {
	if cv.autoMod {
		if err := cv.Modify(i); err != nil {
			return err
		}
	}
	return cv.validator.Struct(i)
}

// Modify to modify/set struct field value according
// to modifier tag. Param `data` should be a pointer.
func (cv customValidator) Modify(data interface{}) error {
	return cv.modifier.Struct(context.Background(), data)
}

type httpController interface {
	Start(host string, port int)
}

type httpCtr struct {
	echo *echo.Echo
}

func NewHttpController() httpController {
	e := echo.New()

	// Middleware

	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	if config.Configuration().Environment == "prod" {
		e.Logger.SetLevel(log.INFO)
	} else {
		e.Logger.SetLevel(log.DEBUG)
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	e.Validator = &customValidator{
		validator: validator.New(),
		modifier:  modifiers.New(),
		autoMod:   true,
	}

	e.Binder = &binder.CustomBinder{}

	// start main cotroller
	return &httpCtr{e}
}

// register cotrollers
func (ctr *httpCtr) Start(host string, port int) {
	book.NewBookController(ctr.echo).Routes("book")
	ctr.echo.Logger.Fatal(ctr.echo.Start(fmt.Sprintf("%s:%d", host, port)))
}
