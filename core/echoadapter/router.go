package echoadapter

import (
	"github.com/ahmadaidin/echoscratch/core"
	"github.com/labstack/echo/v4"
)

type IRouter interface {
	Group(prefix string, m ...echo.MiddlewareFunc) (g IGroup)
}

type IGroup interface {
	Use(middleware ...echo.MiddlewareFunc)

	// CONNECT(path string, h HandlerFunc, m ...echo.MiddlewareFunc)

	// DELETE(path string, h HandlerFunc, m ...echo.MiddlewareFunc)

	GET(path string, h core.HandlerFunc, m ...echo.MiddlewareFunc)

	// HEAD(path string, h HandlerFunc, m ...echo.MiddlewareFunc)

	// OPTIONS(path string, h HandlerFunc, m ...echo.MiddlewareFunc)

	// // PATCH implements `Echo#PATCH()` for sub-routes within the Group.
	// PATCH(path string, h HandlerFunc, m ...echo.MiddlewareFunc)

	// // POST implements `Echo#POST()` for sub-routes within the Group.
	// POST(path string, h HandlerFunc, m ...echo.MiddlewareFunc)

	// // PUT implements `Echo#PUT()` for sub-routes within the Group.
	// PUT(path string, h HandlerFunc, m ...echo.MiddlewareFunc)

	// // TRACE implements `Echo#TRACE()` for sub-routes within the Group.
	// TRACE(path string, h HandlerFunc, m ...echo.MiddlewareFunc)

	// // Any implements `Echo#Any()` for sub-routes within the Group.
	// Any(path string, handler HandlerFunc, middleware ...echo.MiddlewareFunc) []IRouter

	// // Match implements `Echo#Match()` for sub-routes within the Group.
	// Match(methods []string, path string, handler HandlerFunc, middleware ...echo.MiddlewareFunc) []IRouter

	// SubGroup creates a new sub-group with prefix and optional sub-group-level middleware.
	// SubGroup(prefix string, middleware ...echo.MiddlewareFunc) (sg IGroup)

	// // File implements `Echo#File()` for sub-routes within the Group.
	// File(path, file string)

	// RouteNotFound(path string, h HandlerFunc, m ...echo.MiddlewareFunc)

	// Add(method, path string, handler HandlerFunc, middleware ...echo.MiddlewareFunc)
}
