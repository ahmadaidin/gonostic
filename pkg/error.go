package pkg

import (
	"github.com/pkg/errors"
)

const (
	ErrBadParam   = "bad_req_param"
	ErrBadBody    = "bad_req_body"
	ErrUnexpected = "unexpected"
)

type Error struct {
	Code    string `json:"code"`
	Message string `json:"messages,omitempty"`
	Detail  string `json:"detail,omitempty"`
	err     error  `json:"-"`
}

func NewError(code string, err error, messages ...string) (resErr Error) {
	resErr = Error{
		Code: code,
		err:  err,
	}
	if len(messages) > 0 {
		resErr.Message = messages[0]
	}
	resErr.Detail = err.Error()
	return
}

func (err Error) Error() string {
	return errors.Wrap(err.err, err.Message).Error()
}

func (err *Error) RemoveDetail() *Error {
	err.Detail = ""
	return err
}
