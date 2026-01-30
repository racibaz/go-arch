package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/racibaz/go-arch/pkg/helper"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// todo remove this print
		println("Authenticate middleware called")

		authHeader := strings.TrimSpace(c.GetHeader(string(helper.CtxAuthorization)))
		if authHeader == "" || !strings.HasPrefix(strings.ToLower(authHeader), "bearer ") {
			helper.ErrorResponse(
				c,
				"Unauthorized",
				errors.New("Authorization field is empty"),
				http.StatusUnauthorized,
			)
			return
		}

		platform := strings.ToLower(strings.TrimSpace(c.GetHeader(string(helper.CtxPlatform))))
		if platform != helper.PlatformWeb && platform != helper.PlatformMobile {
			helper.ErrorResponse(
				c,
				"Invalid platform",
				errors.New("Invalid platform"),
				http.StatusBadRequest,
			)
			return
		}

		accessToken := strings.TrimSpace(authHeader[7:])

		userId, name, tokenPlatform, err := helper.VerifyJWT(accessToken)
		if err != nil {
			helper.ErrorResponse(
				c,
				"Unauthorized",
				errors.New("Unauthorized"),
				http.StatusUnauthorized,
			)
			return
		}

		// TODO check if refresh token is available in db or cache

		if tokenPlatform != platform {
			helper.ErrorResponse(
				c,
				"Unauthorized",
				errors.New("Unauthorized"),
				http.StatusUnauthorized,
			)
			return
		}

		c.Set(helper.CtxUserID, userId)
		c.Set(helper.CtxUserDisplayName, name)
		c.Set(helper.CtxPlatform, tokenPlatform)

		c.Next()
	}
}
