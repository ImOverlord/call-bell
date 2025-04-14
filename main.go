package main

import (
	"ImOverlord/call-bell/internal/config"
	"ImOverlord/call-bell/internal/env"
	"ImOverlord/call-bell/internal/jokes"
	"ImOverlord/call-bell/internal/logger"
	"ImOverlord/call-bell/internal/metrics"
	"ImOverlord/call-bell/internal/notification"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplates() *Template {
	return &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

type State struct {
	Buttons []config.Buttons
	Message string
}

func main() {
	envConfig := env.Load()

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		RequestIDHandler: func(c echo.Context, s string) {
			c.Set("RequestID", s)
		},
	}))
	e.Use(logger.AddLogger())
	e.Use(logger.LoggerMiddleware())
	e.Static("/", "views")
	e.Renderer = newTemplates()
	metrics.Handler(e)

	var notificationService notification.Notification
	switch envConfig.NotificationType {
	case "mattermost":
		notificationService = notification.NewMattermostNotification(envConfig.NotificationURL)
	case "slack":
		notificationService = notification.NewSlackNotification(envConfig.NotificationURL)
	default:
		log.Fatal().Msg("NotificationType is invalid")
	}

	conf := config.ReadConfig()

	e.GET("/", func(context echo.Context) error {
		return context.Render(http.StatusOK, "index", State{Buttons: conf.Buttons})
	})

	e.POST("/notify", func(context echo.Context) error {
		logger := logger.GetContextLogger(context)
		queryType := context.QueryParam("type")
		msg := "button has been clicked"
		for _, button := range conf.Buttons {
			if button.Type == queryType {
				msg = button.Message
				break
			}
		}
		if err := notificationService.Notify(msg); err != nil {
			logger.Error().Err(err).Msg("Error while creating notification")
			return context.Render(http.StatusInternalServerError, "message", State{Buttons: conf.Buttons, Message: "Error"})
		}
		return context.Render(http.StatusOK, "message", State{Buttons: conf.Buttons, Message: jokes.GetJoke()})
	})
	log.Info().Msgf("Starting server on port %s", envConfig.Port)
	log.Fatal().Err(e.Start(fmt.Sprintf(":%s", envConfig.Port))).Msg("Error while starting HTTP server")
}
