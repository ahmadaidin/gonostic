package http

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/ahmadaidin/echoscratch/config"
	"github.com/ahmadaidin/echoscratch/controller/http/book"
	"github.com/ahmadaidin/echoscratch/core/echoadapter"
	"github.com/ahmadaidin/echoscratch/core/fiberadapter" // we use echo version 4 here
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	_ "github.com/ahmadaidin/echoscratch/docs"
)

// customValidator to validate input request
type customValidator struct {
	validator *validator.Validate
}

// validate using custom validator
func (cv customValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

type HttpHandler interface {
	Listen(port int)
}

type echoHandler struct {
	echo     *echoadapter.Echo
	bookCtrl *book.BookController
}

func (handler *echoHandler) Listen(port int) {
	bookRouter := handler.echo.Group("book")
	bookRouter.GET("", handler.bookCtrl.FindAll)

	handler.echo.Start(fmt.Sprintf(":%d", port))
}

func NewEchoHttpHandler(
	bookCtrl *book.BookController,
) HttpHandler {
	e := echoadapter.NewEcho()

	if config.GetConfig().Environment == "prod" {
		e.Logger.SetLevel(log.INFO)
	} else {
		e.Logger.SetLevel(log.DEBUG)
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	e.Validator = &customValidator{
		validator: validator.New(),
	}

	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return &echoHandler{
		echo:     e,
		bookCtrl: bookCtrl,
	}
}

type fiberHandler struct {
	app      *fiberadapter.Fiber
	bookCtrl *book.BookController
}

func (handler *fiberHandler) Listen(port int) {
	bookRouter := handler.app.Group("book")
	bookRouter.Get("", handler.bookCtrl.FindAll)

	handler.app.Listen(fmt.Sprintf(":%d", port))
}

func NewFiberHttpHandler(
	bookCtrl *book.BookController,
) HttpHandler {
	app := fiberadapter.NewFiber()

	return &fiberHandler{
		app:      app,
		bookCtrl: bookCtrl,
	}
}
