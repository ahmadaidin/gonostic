package core

import "github.com/labstack/echo/v4"

type Router interface {
	Group(prefix string, m ...echo.MiddlewareFunc)
	Use()
}
