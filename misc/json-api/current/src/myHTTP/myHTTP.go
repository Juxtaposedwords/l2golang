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
func (e *HTTPError) GetCode() int {
	return e.Code
}

var (
	Unprocessable = &HTTPError{http.StatusPreconditionFailed, "Invalid insert"}
	NotFoundErr   = &HTTPError{http.StatusNotFound, "Page not found."}
	NotAnIntForID = &HTTPError{http.StatusPreconditionFailed, "A non-integer was provided for an id"}
	WrongMethod   = &HTTPError{http.StatusMethodNotAllowed, "Invalid method used on URI"}
	Created       = &HTTPError{http.StatusCreated, "Sucessfully created"}
	Accepted      = &HTTPError{http.StatusAccepted, "Sucessfully updated"}
	InternalError = &HTTPError{http.StatusInternalServerError, "Something went wrong internally"}
	OK            = &HTTPError{http.StatusOK, "OK"}
)
