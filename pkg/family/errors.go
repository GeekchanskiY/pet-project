package family

import "errors"

var (
	// ErrFamilyCycle occurs when parent or child node contain current node Human
	ErrFamilyCycle = errors.New("family cycle detected")
)
