package middleware

import (
	"github.com/andrewshostak/awesome-service/errs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandle() gin.HandlerFunc {
	return func (c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err == nil {
			return
		}

		if e, ok := err.Err.(errs.Error); ok {
			if e.GetCause() == errs.UserError {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

