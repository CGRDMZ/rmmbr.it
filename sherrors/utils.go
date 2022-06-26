package sherrors

import (
	"github.com/CGRDMZ/rmmbrit-api/config"
)

func HandleError(err error) *CustomError {
	switch e := err.(type) {
	case CustomError:
		return &e
	case *CustomError:
		return e
	default:
		if config.Conf.Env == "PROD" {
			return ErrInternalServer
		}
		
		intErr := InternalError(e)
		return &intErr
	}
}
