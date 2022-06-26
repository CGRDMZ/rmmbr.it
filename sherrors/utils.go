package sherrors

import (
	"github.com/CGRDMZ/rmmbrit-api/config"
)

func HandleError(err error) *CustomError {
	switch e := err.(type) {

	case CustomError:
		if config.Conf.Env == "PROD" {
			e.InnerError = nil
		}
		return &e
	case *CustomError:
		if config.Conf.Env == "PROD" {
			e.InnerError = nil
		}
		return e
	default:
		if config.Conf.Env == "PROD" {
			return ErrInternalServer
		}

		intErr := InternalError(e)
		return &intErr
	}
}
