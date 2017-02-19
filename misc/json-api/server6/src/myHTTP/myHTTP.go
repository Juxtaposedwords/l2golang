package myHTTP

import (
	"net/http"
)

type HTTPError struct {
	Code int
	Err  string
}

func (e *HTTPError) Error() string {
	return e.Err
}

var (
	Unprocessable = &HTTPError{http.StatusPreconditionFailed, "Invalid insert"}
	NotFoundErr   = &HTTPError{http.StatusNotFound, "Page not found."}
)
