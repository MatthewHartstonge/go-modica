package modica

import "errors"

// Modica Generic API errors
var (
	// ErrNotFound provides a generic 404 not found error
	ErrNotFound = errors.New("not found")
)
