package jwt

import (
	"crypto/rand"
	"time"

	"github.com/SpeedCrash100/go-login-service/pkg/consts"
	m "github.com/SpeedCrash100/go-login-service/pkg/models"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

const JWT_SECRET_DEFAULT_SIZE int = 512
const IDENTITY_KEY = "username"

func GetJWTSecret() []byte {

	jwt_secret := []byte(viper.GetString(consts.CONFIG_JWT_SECRET))
	// Generate random
	if len(jwt_secret) == 0 {
		jwt_secret = generateJWTKey()
	}

	return []byte(jwt_secret)
}

func generateJWTKey() []byte {
	new_key := make([]byte, JWT_SECRET_DEFAULT_SIZE)

	_, err := rand.Read(new_key)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to generate jwt key")
	}

	return new_key
}

func GetJWTMiddleware() (*jwt.GinJWTMiddleware, error) {

	jwtMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "deka.space",
		Key:             GetJWTSecret(),
		Timeout:         time.Minute,
		MaxRefresh:      time.Minute,
		IdentityKey:     IDENTITY_KEY,
		PayloadFunc:     payloadFunc,
		IdentityHandler: identityHandler,
		Authenticator:   authenticator,
		Authorizator:    authorization,
		Unauthorized:    unauthorized,
	})

	if err != nil {
		return nil, err
	}

	err = jwtMiddleware.MiddlewareInit()
	return jwtMiddleware, err
}

func payloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*m.UserClaims); ok {
		return jwt.MapClaims{
			IDENTITY_KEY: v.Username,
		}
	}
	return jwt.MapClaims{}
}

func identityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &m.UserClaims{
		Username: claims[IDENTITY_KEY].(string),
	}
}

func authenticator(c *gin.Context) (interface{}, error) {
	var login m.UserLoginInfo
	if err := c.ShouldBind(&login); err != nil {
		return "", jwt.ErrMissingLoginValues
	}

	// TODO! Check username/password in DB

	return &m.UserClaims{Username: login.Username}, nil
}

func authorization(data interface{}, c *gin.Context) bool {
	if _, ok := data.(*m.UserClaims); ok {
		//TODO! Check rights
		return true
	}
	return false
}

func unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}
