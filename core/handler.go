package core

type HandlerFunc func(c Context) error

type Runner interface {
	Start(address string) error
}
