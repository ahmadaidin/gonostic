package adapter

import (
	"github.com/labstack/echo/v4"
)

func NewHTTPError(code int, messages ...any) error {
	return echo.NewHTTPError(code, messages...)
}
