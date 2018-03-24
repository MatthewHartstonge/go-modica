package modica

import "errors"

// Modica Generic API errors
var (
	// ErrUnauthorized provides a generic 401 unauthorized error
	ErrUnauthorized = errors.New("not authorized - valid authentication credentials for the target resource are required")

	// ErrNotFound provides a generic 404 not found error
	ErrNotFound = errors.New("not found")
)
