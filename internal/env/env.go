package env

import (
	"github.com/caarlos0/env/v6"
	"github.com/rs/zerolog/log"
)

type Env struct {
	NotificationURL  string `env:"NOTIFICATION_URL,required"`
	NotificationType string `env:"NOTIFICATION_TYPE,required"`
	Port             string `env:"PORT" envDefault:"9000"`
}

func Load() *Env {
	cfg := Env{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal().Err(err).Msg("Failed to parse environment variables")
	}
	return &cfg
}
