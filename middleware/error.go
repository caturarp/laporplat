package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/caturarp/laporplat/apperr"
	"github.com/caturarp/laporplat/dto"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.Last()
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				c.AbortWithStatusJSON(http.StatusGatewayTimeout, dto.Response{Message: "request timeout"})
				return
			}
			switch e := err.Err.(type) {
			case *apperr.CustomError:
				c.AbortWithStatusJSON(e.Code, e.ConvertToErrorResponse())
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
				return
			}
		}
	}
}
