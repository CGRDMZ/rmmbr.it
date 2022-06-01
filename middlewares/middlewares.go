package middlewares

import (
	"log"
	"net/http"

	"github.com/CGRDMZ/rmmbrit-api/errors"
	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) <= 0 {
		return
	}

	var err = c.Errors[0]

	if ce, ok := err.Err.(*errors.CustomError); ok {

		if ce.StatusCode == http.StatusNotFound {
			c.HTML(http.StatusNotFound, "not-found.html", gin.H{})
			log.Print("hello")
			return
		}

		c.JSON(ce.StatusCode, ce)
		return
	}

	errInternal := errors.NewCustomError("Something unexpected happened. Hmm...", err.Error(), 500, nil)
	c.JSON(500, errInternal)

}
