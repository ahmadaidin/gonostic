package http

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/ahmadaidin/gonostic/config"
	"github.com/ahmadaidin/gonostic/controller/http/bookctrl"
	"github.com/ahmadaidin/gonostic/core/echoadapter"
	"github.com/ahmadaidin/gonostic/core/fiberadapter" // we use echo version 4 here
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	_ "github.com/ahmadaidin/gonostic/docs"
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
	echo         *echoadapter.Echo
	bookctrlCtrl *bookctrl.BookController
}

func (handler *echoHandler) Listen(port int) {
	bookctrlRouter := handler.echo.Group("bookctrl")
	bookctrlRouter.GET("", handler.bookctrlCtrl.FindAll)

	handler.echo.Start(fmt.Sprintf(":%d", port))
}

func NewEchoHttpHandler(
	bookctrlCtrl *bookctrl.BookController,
) HttpHandler {
	e := echoadapter.NewEcho()

	if config.GetConfig().IsProdEnv() {
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
		echo:         e,
		bookctrlCtrl: bookctrlCtrl,
	}
}

type fiberHandler struct {
	app          *fiberadapter.Fiber
	bookctrlCtrl *bookctrl.BookController
}

func (handler *fiberHandler) Listen(port int) {
	bookctrlRouter := handler.app.Group("bookctrl")
	bookctrlRouter.Get("", handler.bookctrlCtrl.FindAll)

	handler.app.Listen(fmt.Sprintf(":%d", port))
}

func NewFiberHttpHandler(
	bookctrlCtrl *bookctrl.BookController,
) HttpHandler {
	app := fiberadapter.NewFiber()

	return &fiberHandler{
		app:          app,
		bookctrlCtrl: bookctrlCtrl,
	}
}
