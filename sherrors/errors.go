package sherrors

import (
	"fmt"
	"net/http"
	"time"
)

type CustomError struct {
	Title      string    `json:"title"`
	Detail     string    `json:"detail"`
	StatusCode int       `json:"statusCode"`
	OccuredAt  time.Time `json:"occuredAt"`
	InnerError error    `json:"error,omitempty"`
}

func NewCustomError(title, detail string, code int, err error) *CustomError {
	return &CustomError{
		Title:      title,
		Detail:     detail,
		InnerError: err,
		StatusCode: code,
		OccuredAt:  time.Now(),
	}
}

func (ce CustomError) Error() string {
	return fmt.Sprintf("[Error %d][%s] Title:%s, Detail: %s", ce.StatusCode, ce.OccuredAt.String(), ce.Title, ce.Detail)
}

func (ce CustomError) String() string {
	return fmt.Sprintf("[Error %d][%s] Title:%s, Detail: %s", ce.StatusCode, ce.OccuredAt.String(), ce.Title, ce.Detail)
}

var (
	ErrCommonNotFound = NewCustomError("the requested resource is not found.", "there were no resorce in the database with the given parameters.", http.StatusNotFound, nil)
	ErrLoginFailed    = NewCustomError("login failed.", "the given credentials are not valid.", http.StatusUnauthorized, nil)
	ErrInternalServer = NewCustomError("Something unexpected happened", "This is a server error, contact to the owners of this site if your problem persist.", 500, nil)
)

func NotFoundErr(resourceName, id string) CustomError {
	return *NewCustomError(fmt.Sprintf("the requested '%s' with id '%s' is not found.", resourceName, id), "there were no resorce in the database with the given parameters.", http.StatusNotFound, nil)
}

func AlreadyExistsErr(resourceName string) CustomError {
	return *NewCustomError(fmt.Sprintf("the requested '%s' already exists.", resourceName), fmt.Sprintf("there is already a record with this name. '%s' should be unique.", resourceName), http.StatusConflict, nil)
}

func InternalError(err error) CustomError {
	return *NewCustomError("Something unexpected happened", "This is a server error, contact to the owners of this site if your problem persist.", 500, err)
}

// func InternalErrorWithInnerError(err error) CustomError {
// 	errInternal := *ErrInternalServer

// 	*errInternal.InnerError = err

// 	return errInternal
// }
