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
	NotFoundErr = &HTTPError{http.StatusNotFound, "Page not found."}
)
