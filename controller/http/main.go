package http

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/ahmadaidin/echoscratch/config"
	"github.com/ahmadaidin/echoscratch/controller/http/book"
	"github.com/ahmadaidin/echoscratch/domain/repository/mongorepo"
	"github.com/ahmadaidin/echoscratch/infra"
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
	echo *echo.Echo
}

func NewHttpHandler() HttpHandler {
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
	return &httpHandler{e}
}

func (handler *httpHandler) Listen(port int) {
	router := handler.echo
	runner := handler.echo

	mongoConnection := infra.NewMongoConnection("")

	bookRepo := mongorepo.NewBookRepository(mongoConnection)

	bookCtrl := book.NewBookController(
		bookRepo,
	)

	rg := router.Group("book")
	rg.GET("", bookCtrl.FindAll)

	runner.Start(fmt.Sprintf(":%d", port))
}
