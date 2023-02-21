package core

type MiddlewareFunc func(next HandlerFunc) HandlerFunc
