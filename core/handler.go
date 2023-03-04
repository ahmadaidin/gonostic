package core

type HandlerFunc func(c Context) error

type MiddlewareFunc func(next HandlerFunc) HandlerFunc

type Runner interface {
	Start(address string) error
}
