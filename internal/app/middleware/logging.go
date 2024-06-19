package middleware

import (
	"time"

	"bytes"
	logger "github.com/gcarrenho/common-libs/pkg/logging"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		if raw != "" {
			path = path + "?" + raw
		}

		latency := time.Since(startTime).String()
		statusCode := c.Writer.Status()
		method := c.Request.Method
		clientIP := c.ClientIP()
		contentType := c.Request.Header.Get("Content-type")
		requestID := c.Request.Header.Get("X-Request-Id")

		logging := logger.NewLogging("HTTP request")
		logging.HttpMethod = &method
		logging.Path = &path
		logging.StatusCode = &statusCode
		logging.RequestID = &requestID
		logging.RemoteIP = &clientIP
		logging.ContentType = &contentType
		logging.Latency = &latency

		c.Next()

		// Create a buffer and logger
		var buf bytes.Buffer
		logger := zerolog.New(&buf).With().Logger()

		// Log the message
		logger.Log().Object("log", logging).Send()

		// Print the log output
		trimmedLog := buf.String()
		trimmedLog = trimmedLog[:len(trimmedLog)-1] // Remove the extra newline character

		log.Info().RawJSON("ginlog", []byte(trimmedLog)).Send()

		// Empty the buffer
		buf.Reset()
	}
}
