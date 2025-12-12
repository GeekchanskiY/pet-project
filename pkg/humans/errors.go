package humans

import "errors"

var (
	ErrNotEnoughArguments = errors.New("not enough arguments")
	ErrInvalidArguments   = errors.New("invalid arguments")
)
