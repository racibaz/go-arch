package helper

import (
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	CtxUserID          string = "userId"
	CtxUserDisplayName string = "name"
	CtxPlatform        string = "X-Platform"
	CtxAuthorization   string = "Authorization"
	PlatformWeb               = "web"
	PlatformMobile            = "mobile"
)

func Platform(c *gin.Context) string {
	return strings.ToLower(
		strings.TrimSpace(c.GetHeader(string(CtxPlatform))),
	)
}
