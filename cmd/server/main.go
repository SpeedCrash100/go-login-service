package main

import (
	_ "github.com/SpeedCrash100/go-login-service/internal/config"
	"github.com/SpeedCrash100/go-login-service/internal/database"
	"github.com/SpeedCrash100/go-login-service/internal/database/query"
	my_jwt "github.com/SpeedCrash100/go-login-service/internal/jwt"
	"github.com/SpeedCrash100/go-login-service/internal/routes"
	"github.com/SpeedCrash100/go-login-service/pkg/consts"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func initRoutes(r *gin.Engine, jwtMiddleware *jwt.GinJWTMiddleware) {
	r.POST("/login", jwtMiddleware.LoginHandler)
	r.POST("/register", routes.Register)

	auth := r.Group("/auth")
	{
		auth.Use(jwtMiddleware.MiddlewareFunc())
		auth.GET("/refresh_token", jwtMiddleware.RefreshHandler)
		auth.GET("/profile", routes.Profile)

	}
}

func main() {

	// Set Logging level
	if viper.GetBool(consts.CONFIG_IS_DEBUG) {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	// Prepare engine
	e := gin.New()
	e.Use(logger.SetLogger())
	e.Use(gin.Recovery())

	// Prepare JWT
	jwtMiddleware, err := my_jwt.GetJWTMiddleware()
	if err != nil {
		log.Fatal().Err(err).Msg("jwt init")
	}

	// Prepare DB
	db := database.NewUserDB()
	database.Migrate(db)
	query.SetDefault(db)

	initRoutes(e, jwtMiddleware)

	// Start
	address := viper.GetString(consts.CONFIG_IP) + ":" + viper.GetString(consts.CONFIG_PORT)
	e.Run(address)

}
