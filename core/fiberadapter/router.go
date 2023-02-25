package fiberadapter

import (
	"github.com/ahmadaidin/echoscratch/core"
)

type IRouter interface {
	Use(args ...interface{}) IRouter

	Get(path string, handlers ...core.HandlerFunc) IRouter
	Head(path string, handlers ...core.HandlerFunc) IRouter
	Post(path string, handlers ...core.HandlerFunc) IRouter
	Put(path string, handlers ...core.HandlerFunc) IRouter
	Delete(path string, handlers ...core.HandlerFunc) IRouter
	Connect(path string, handlers ...core.HandlerFunc) IRouter
	Options(path string, handlers ...core.HandlerFunc) IRouter
	Trace(path string, handlers ...core.HandlerFunc) IRouter
	Patch(path string, handlers ...core.HandlerFunc) IRouter

	Add(method, path string, handlers ...core.HandlerFunc) IRouter
	// Static(prefix, root string, config ...Static) IRouter
	All(path string, handlers ...core.HandlerFunc) IRouter

	Group(prefix string, m ...core.HandlerFunc) IRouter
}
