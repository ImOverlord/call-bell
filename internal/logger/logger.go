package logger

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

func GetContextLogger(c echo.Context) zerolog.Logger {
	logger, ok := c.Get("logger").(zerolog.Logger)
	if !ok {
		logger = zerolog.New(os.Stdout)
		logger.Warn().Msg("Failed to fetch logger from Context")
	}

	requestID, ok := c.Get("RequestID").(string)
	if ok {
		logger = logger.With().Str("RequestID", requestID).Logger()
	}
	return logger
}

func AddLogger() echo.MiddlewareFunc {
	logger := zerolog.New(os.Stdout)
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("logger", logger)
			return next(c)
		}
	}
}

func Middleware() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger := GetContextLogger(c)
			logger.Info().
				Str("URI", v.URI).
				Int("status", v.Status).
				Str("IP", v.RemoteIP).
				Msg("request")

			return nil
		},
	})
}
