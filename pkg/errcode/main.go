package errcode

import "errors"

var (
	ErrBadParam = errors.New("bad query param")
	ErrBadBody  = errors.New("bad request body")
)
