package config

import (
	"github.com/rs/zerolog/log"

	"github.com/spf13/viper"
)

type Buttons struct {
	Type    string `json:"type"`
	Text    string `json:"text"`
	Message string `json:"message"`
}

type Config = struct {
	Buttons []Buttons `json:"buttons"`
}

func ReadConfig() *Config {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Err(err).Msg("Config error")
	}
	conf := &Config{}
	err := viper.Unmarshal(conf)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to unmarshal config file")
	}
	return conf
}
