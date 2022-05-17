package errors

import (
	"fmt"
	"net/http"
	"time"
)

type CustomError struct {
	Title      string       `json:"title"`
	Detail     string       `json:"detail"`
	StatusCode int          `json:"statusCode"`
	OccuredAt  time.Time    `json:"occuredAt"`
	Payload    interface{} `json:"payload,omitempty"`
}

func NewCustomError(title, detail string, code int, payload interface{}) *CustomError {
	return &CustomError{
		Title:     title,
		Detail:    detail,
		Payload:   payload,
		StatusCode: code,
		OccuredAt: time.Now(),
	}
}

func (ce CustomError) Error() string {
	return ce.Title
}

var (
	ErrCommonNotFound = NewCustomError("the requested resource is not found.", "there were no resorce in the database with the given parameters.", http.StatusNotFound, nil)
	ErrLoginFailed = NewCustomError("login failed.", "the given credentials are not valid.", http.StatusUnauthorized, nil)
)

func NotFoundErr(resourceName, id string) *CustomError {
	return NewCustomError(fmt.Sprintf("the requested '%s' witlh id '%s' is not found.", resourceName, id), "there were no resorce in the database with the given parameters.", http.StatusNotFound, nil)
}
