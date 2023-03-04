package errors

import (
	"strings"

	"github.com/pkg/errors"
)

const (
	ErrBadParam   = "bad_req_param"
	ErrBadBody    = "bad_req_body"
	ErrUnexpected = "unexpected"
	ErrUnknown    = "unknown"
)

type Error struct {
	Code    string   `json:"code"`
	Message string   `json:"message,omitempty"`
	Details []string `json:"details,omitempty"`
	err     error    `json:"-"`
}

type multipleResultUnwrappable interface {
	Unwrap() []error
}

type singleResultUnwrappable interface {
	Unwrap() error
}

func NewError(code string, err error, messages ...string) (resErr Error) {
	resErr = Error{
		Code: code,
		err:  err,
	}
	if len(messages) > 0 {
		resErr.Message = messages[0]
	} else {
		join, ok := err.(multipleResultUnwrappable)
		if ok {
			e := join.Unwrap()
			if len(e) > 0 {
				realerr, ok := e[0].(singleResultUnwrappable)
				if ok {
					resErr.Message = realerr.Unwrap().Error()
				} else {
					resErr.Message = ErrUnknown
				}
			}
		} else {
			join, ok := err.(singleResultUnwrappable)
			if ok {
				e := join.Unwrap()
				if e != nil {
					resErr.Message = e.Error()
				}
			} else {
				resErr.Message = ErrUnknown
			}
		}
	}
	details := []string{}
	details = append(details, messages...)
	details = append(details, strings.Split(err.Error(), "\n")...)
	resErr.Details = details
	return
}

func (err Error) Error() string {
	return errors.Wrap(err.err, err.Message).Error()
}

func (err *Error) RemoveDetails() *Error {
	err.Details = nil
	return err
}
