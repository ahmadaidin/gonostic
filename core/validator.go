package core

type Validator interface {
	Validate(any) error
}
