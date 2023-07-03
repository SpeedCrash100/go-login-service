package config

import (
	"os"

	"github.com/SpeedCrash100/go-login-service/pkg/consts"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func init() {
	// Start-up log configuration
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	viper.SetDefault(consts.CONFIG_PORT, 8080)
	viper.SetDefault(consts.CONFIG_IP, "0.0.0.0")
	viper.SetDefault(consts.CONFIG_IS_DEBUG, true)
	viper.SetDefault(consts.CONFIG_USER_DB_CONN_STRING, "sqlite.db")

	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Warn().Err(err).Msg("failed to find config file. using default value")
		} else {
			log.Fatal().Err(err).Msg("failed to open config file")
		}
	}
}
