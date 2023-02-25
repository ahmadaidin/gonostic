package echoadapter

import (
	"github.com/ahmadaidin/echoscratch/core"
	"github.com/labstack/echo/v4"
)

type (
	Echo struct {
		*echo.Echo
	}
	Group struct {
		*echo.Group
	}
	Context struct {
		echo.Context
	}
)

func NewEcho(e *echo.Echo) *Echo {
	return &Echo{e}
}

func (e *Echo) Group(prefix string, m ...echo.MiddlewareFunc) IGroup {
	g := e.Echo.Group(prefix, m...)
	return &Group{Group: g}
}

func (g *Group) Use(m ...echo.MiddlewareFunc) {

	g.Group.Use(m...)
}

func (g *Group) CONNECT(path string, h core.HandlerFunc, m ...echo.MiddlewareFunc) {
	hdl := func(c echo.Context) error {
		return h(&Context{c})
	}

	g.Group.CONNECT(path, hdl, m...)
}

func (g *Group) DELETE(path string, h core.HandlerFunc, m ...echo.MiddlewareFunc) {
	hdl := func(c echo.Context) error {
		return h(&Context{c})
	}

	g.Group.DELETE(path, hdl, m...)
}

func (g *Group) GET(path string, h core.HandlerFunc, m ...echo.MiddlewareFunc) {
	hdl := func(c echo.Context) error {
		return h(&Context{c})
	}

	g.Group.GET(path, hdl, m...)
}

func (g *Group) HEAD(path string, h core.HandlerFunc, m ...echo.MiddlewareFunc) {
	hdl := func(c echo.Context) error {
		return h(&Context{c})
	}

	g.Group.HEAD(path, hdl, m...)
}

func (g *Group) OPTIONS(path string, h core.HandlerFunc, m ...echo.MiddlewareFunc) {
	hdl := func(c echo.Context) error {
		return h(&Context{c})
	}

	g.Group.OPTIONS(path, hdl, m...)
}

func (g *Group) PATCH(path string, h core.HandlerFunc, m ...echo.MiddlewareFunc) {
	hdl := func(c echo.Context) error {
		return h(&Context{c})
	}

	g.Group.PATCH(path, hdl, m...)
}

func (g *Group) POST(path string, h core.HandlerFunc, m ...echo.MiddlewareFunc) {
	hdl := func(c echo.Context) error {
		return h(&Context{c})
	}

	g.Group.POST(path, hdl, m...)
}

func (g *Group) PUT(path string, h core.HandlerFunc, m ...echo.MiddlewareFunc) {
	hdl := func(c echo.Context) error {
		return h(&Context{c})
	}

	g.Group.PUT(path, hdl, m...)
}

func (g *Group) TRACE(path string, h core.HandlerFunc, m ...echo.MiddlewareFunc) {
	hdl := func(c echo.Context) error {
		return h(&Context{c})
	}

	g.Group.TRACE(path, hdl, m...)
}

// func (g *Group) Any(path string, handler core.HandlerFunc, middleware ...echo.MiddlewareFunc) []core.IRouter {
// 	return g.Group.Any(path, handler, middleware...)
// }

// func (g *Group) Match(methods []string, path string, handler core.HandlerFunc, middleware ...echo.MiddlewareFunc) []core.IRouter {
// 	return g.Group.Match(methods, path, handler, middleware...)
// }

// Group creates a new sub-group with prefix and optional sub-group-level middleware.
func (g *Group) SubGroup(prefix string, m ...echo.MiddlewareFunc) (sg IGroup) {

	_sg := g.Group.Group(prefix, m...)
	sg = &Group{Group: _sg}
	return sg
}

// // File implements `Echo#File()` for sub-routes within the Group.
// func (g *Group) File(path, file string) {

// }

// // RouteNotFound implements `Echo#RouteNotFound()` for sub-routes within the Group.
// //
// // Example: `g.Group.RouteNotFound("/*", func(c echo.Context) error { return c.NoContent(http.StatusNotFound) })`
// func (g *Group) RouteNotFound(path string, h core.HandlerFunc, m ...echo.MiddlewareFunc) {

// }

// Add implements `Echo#Add()` for sub-routes within the Group.
// func (g *Group) Add(method, path string, handler core.HandlerFunc, middleware ...echo.MiddlewareFunc) {
// 	g.Group.Add(method, path, handler, middleware...)
// }

func (context *Context) SendJson(code int, i interface{}) error {
	return context.Context.JSON(code, i)
}

func (context *Context) QueryParser(i interface{}) error {
	return context.Context.Bind(i)
}
