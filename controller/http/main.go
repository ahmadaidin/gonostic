package http

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/ahmadaidin/echoscratch/config"
	"github.com/ahmadaidin/echoscratch/controller/http/book"
	"github.com/ahmadaidin/echoscratch/core"
	"github.com/ahmadaidin/echoscratch/core/adapter"
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
	runner   core.Runner
	router   core.Router
	bookCtrl *book.BookController
}

func NewEchoHttpHandler(
	bookCtrl *book.BookController,
) HttpHandler {
	e := echo.New()

	// Middleware

	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	if config.GetConfig().Environment == "prod" {
		e.Logger.SetLevel(log.INFO)
	} else {
		e.Logger.SetLevel(log.DEBUG)
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	e.Validator = &customValidator{
		validator: validator.New(),
	}
	wrappedEcho := adapter.NewEcho(e)

	// start main cotroller
	return &httpHandler{
		router:   wrappedEcho,
		runner:   e,
		bookCtrl: bookCtrl,
	}
}

func (handler *httpHandler) Listen(port int) {
	bookRouter := handler.router.Group("book")
	bookRouter.GET("", handler.bookCtrl.FindAll)

	handler.runner.Start(fmt.Sprintf(":%d", port))
}
