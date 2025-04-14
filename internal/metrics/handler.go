package metrics

import (
	"net/http"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
)

func Handler(e *echo.Echo) {
	e.Use(echoprometheus.NewMiddleware(""))

	e.GET("/health", func(context echo.Context) error {
		return context.String(http.StatusOK, "Healthy")
	})
	e.GET("/metrics", echoprometheus.NewHandler())
}
