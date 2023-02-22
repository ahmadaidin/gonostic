package http

import (
	"fmt"

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
}

// validate using custom validator
func (cv customValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

type HttpHandler interface {
	Listen(port int)
}

type httpHandler struct {
	runner   *echo.Echo
	router   *echo.Echo
	bookCtrl *book.BookController
}

func NewHttpHandler(
	bookCtrl *book.BookController,
) HttpHandler {
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
	}

	e.Binder = &binder.CustomBinder{}

	// start main cotroller
	return &httpHandler{
		router:   e,
		runner:   e,
		bookCtrl: bookCtrl,
	}
}

func (handler *httpHandler) Listen(port int) {
	rg := handler.router.Group("book")
	rg.GET("", handler.bookCtrl.FindAll)

	handler.runner.Start(fmt.Sprintf(":%d", port))
}
