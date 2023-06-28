package main

import (
	"net/http"
	"os"

	"github.com/SpeedCrash100/go-login-service/pkg/consts"
	"github.com/gin-contrib/logger"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func initConfig() {
	viper.SetDefault(consts.CONFIG_PORT, 8080)
	viper.SetDefault(consts.CONFIG_IP, "0.0.0.0")
	viper.SetDefault(consts.CONFIG_IS_DEBUG, true)

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

func initRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}

func main() {
	// Start-up log configuration
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	// Configs
	initConfig()

	// Set Logging level
	if viper.GetBool(consts.CONFIG_IS_DEBUG) {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	// Prepare engine
	e := gin.New()
	e.Use(logger.SetLogger())
	e.Use(gin.Recovery())

	initRoutes(e)

	// Start
	address := viper.GetString(consts.CONFIG_IP) + ":" + viper.GetString(consts.CONFIG_PORT)
	e.Run(address)

}
