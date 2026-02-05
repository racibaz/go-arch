package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/racibaz/go-arch/pkg/helper"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

const (
	InValidPlatform = "invalid platform provided"
)

func Platform() gin.HandlerFunc {
	return func(c *gin.Context) {
		platform := helper.Platform(c)

		if platform != helper.PlatformWeb && platform != helper.PlatformMobile {

			if span := trace.SpanFromContext(c); span != nil {
				span.SetAttributes(attribute.String("error", InValidPlatform))
				span.SetStatus(codes.Error, InValidPlatform)
				span.RecordError(errors.New(InValidPlatform))
			}

			helper.ErrorResponse(
				c,
				InValidPlatform,
				errors.New(InValidPlatform),
				http.StatusBadRequest,
			)
			return
		}
		c.Next()
	}
}
