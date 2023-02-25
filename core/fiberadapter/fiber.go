package fiberadapter

import (
	"encoding/json"
	"errors"

	"github.com/ahmadaidin/echoscratch/core"
	"github.com/gofiber/fiber/v2"
)

type (
	Fiber struct {
		*fiber.App
	}
	Context struct {
		*fiber.Ctx
	}
	Router struct {
		fiber.Router
	}
)

func NewFiber(e *fiber.App) *Fiber {
	return &Fiber{e}
}

func (e *Fiber) Group(prefix string, h ...fiber.Handler) IRouter {
	r := e.App.Group(prefix, h...)
	return &Router{Router: r}
}

func convertHandlers(handlers ...core.HandlerFunc) []fiber.Handler {
	var h []fiber.Handler
	for _, handler := range handlers {
		h = append(h, func(c *fiber.Ctx) error {
			return handler(&Context{Ctx: c})
		})
	}
	return h
}

func (r *Router) Use(args ...interface{}) IRouter {
	return r
}

func (r *Router) Get(path string, handlers ...core.HandlerFunc) IRouter {
	hdl := convertHandlers(handlers...)
	router := r.Router.Get(path, hdl...)
	return &Router{Router: router}
}

func (r *Router) Head(path string, handlers ...core.HandlerFunc) IRouter {
	hdl := convertHandlers(handlers...)
	router := r.Router.Head(path, hdl...)
	return &Router{Router: router}
}

func (r *Router) Post(path string, handlers ...core.HandlerFunc) IRouter {
	hdl := convertHandlers(handlers...)
	router := r.Router.Post(path, hdl...)
	return &Router{Router: router}
}

func (r *Router) Put(path string, handlers ...core.HandlerFunc) IRouter {
	hdl := convertHandlers(handlers...)
	router := r.Router.Put(path, hdl...)
	return &Router{Router: router}
}

func (r *Router) Delete(path string, handlers ...core.HandlerFunc) IRouter {
	hdl := convertHandlers(handlers...)
	router := r.Router.Delete(path, hdl...)
	return &Router{Router: router}
}

func (r *Router) Connect(path string, handlers ...core.HandlerFunc) IRouter {
	hdl := convertHandlers(handlers...)
	router := r.Router.Connect(path, hdl...)
	return &Router{Router: router}
}

func (r *Router) Options(path string, handlers ...core.HandlerFunc) IRouter {
	hdl := convertHandlers(handlers...)
	router := r.Router.Options(path, hdl...)
	return &Router{Router: router}
}

func (r *Router) Trace(path string, handlers ...core.HandlerFunc) IRouter {
	hdl := convertHandlers(handlers...)
	router := r.Router.Trace(path, hdl...)
	return &Router{Router: router}
}

func (r *Router) Patch(path string, handlers ...core.HandlerFunc) IRouter {
	hdl := convertHandlers(handlers...)
	router := r.Router.Patch(path, hdl...)
	return &Router{Router: router}
}

func (r *Router) Add(method string, path string, handlers ...core.HandlerFunc) IRouter {
	hdl := convertHandlers(handlers...)
	router := r.Router.Add(method, path, hdl...)
	return &Router{Router: router}
}

// Static(prefix, root string, config ...Static)
func (r *Router) All(path string, handlers ...core.HandlerFunc) IRouter {
	hdl := convertHandlers(handlers...)
	router := r.Router.All(path, hdl...)
	return &Router{Router: router}
}

func (r *Router) Group(prefix string, handlers ...core.HandlerFunc) IRouter {
	hdl := convertHandlers(handlers...)
	router := r.Router.All(prefix, hdl...)
	return &Router{Router: router}
}

func (ctx *Context) Bind(i interface{}) error {
	m, ok := (i).(fiber.Map)
	if !ok {
		return errors.New("cannot bind to fiber.Map type")
	}
	return ctx.Ctx.Bind(m)
}

func (ctx *Context) SendJson(code int, i interface{}) error {
	return ctx.Ctx.Status(code).JSON(i)
}

func (context *Context) HttpError(code int, messages ...any) error {
	if len(messages) > 0 {
		printedMsg := messages[0]
		_, err := json.Marshal(printedMsg)
		if err == nil {
			context.Ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
			return context.Ctx.Status(code).JSON(printedMsg)
		} else {
			msg, ok := printedMsg.(error)
			if ok {
				return fiber.NewError(code, msg.Error())
			} else {
				msg, ok := printedMsg.(string)
				if ok {
					return fiber.NewError(code, msg)
				} else {
					msg = "unable to convert error message"
					return fiber.NewError(code, msg)
				}
			}
		}
	} else {
		return fiber.NewError(code)
	}
}
