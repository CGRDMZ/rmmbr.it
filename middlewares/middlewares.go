package middlewares

import (
	"log"
	"net/http"

	"github.com/CGRDMZ/rmmbrit-api/sherrors"
	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) <= 0 {
		return
	}

	var err = c.Errors[0]

	ce := sherrors.HandleError(err.Err)

	if ce.StatusCode == http.StatusNotFound {
		c.HTML(http.StatusNotFound, "not-found.html", gin.H{})
		return
	}

	if ce.StatusCode == http.StatusInternalServerError {
		c.HTML(http.StatusInternalServerError, "internal.html", gin.H{"error": ce})
		return
	}

	log.Println(ce)
	c.JSON(ce.StatusCode, ce)
}
