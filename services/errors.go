package services

import (
	"github.com/CGRDMZ/rmmbrit-api/sherrors"
	"net/http"
)

var (
	ErrInvalidUrlFormat = func(inner *error) sherrors.CustomError {
		return *sherrors.NewCustomError("Problem Parsing the Url", "the given string is not a valid URL.", http.StatusBadRequest, *inner)
	}
)
