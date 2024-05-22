package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/caturarp/laporplat/apperr"
	"github.com/caturarp/laporplat/dto"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthorizeHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if gin.Mode() == gin.TestMode {
			return
		}
		whitelistPaths := []string{
			"/auth/register",
			"/auth/login",
			"/users/reset_password",
		}

		for _, path := range whitelistPaths {
			if ctx.Request.URL.Path == path {
				ctx.Next()
				return
			}
		}
		if strings.HasPrefix(ctx.Request.URL.Path, "/auth/verify") {
			ctx.Next()
			return
		}

		var response dto.Response

		header := ctx.Request.Header["Authorization"]
		if len(header) == 0 {
			response.Message = apperr.ErrBearerTokenNotExist.Error()
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		splittedHeader := strings.Split(header[0], " ")
		if len(splittedHeader) != 2 {
			response.Message = apperr.ErrUnauthorized.Error()
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claims := &dto.JWTClaims{}

		token, err := jwt.ParseWithClaims(splittedHeader[1], claims, func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, apperr.ErrWrongCredentials
			}

			return []byte(os.Getenv("API_SECRET")), nil
		})
		if err != nil {
			response.Message = err.Error()
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		_, ok := token.Claims.(*dto.JWTClaims)
		if !ok || !token.Valid {
			response.Message = apperr.ErrUnauthorized.Error()
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		newContext := context.WithValue(ctx.Request.Context(), "user_id", claims.UserID)
		newContext = context.WithValue(newContext, "personal_id", claims.PersonalID)
		ctx.Request = ctx.Request.WithContext(newContext)
		ctx.Next()
	}
}
