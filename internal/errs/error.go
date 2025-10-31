package errs

import "errors"

var (
	ErrNotFound      = errors.New("data not found")
	ErrDuplicateData = errors.New("duplicate data")
)
