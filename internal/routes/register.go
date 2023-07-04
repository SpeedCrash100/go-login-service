package routes

import (
	"net/http"

	"github.com/SpeedCrash100/go-login-service/internal/database/query"
	"github.com/SpeedCrash100/go-login-service/internal/password"
	"github.com/SpeedCrash100/go-login-service/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {
	var user_info models.UserRegisterInfo
	if err := c.ShouldBind(&user_info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	password_hash, err := password.HashPassword(user_info.Password)
	if err != nil {
		log.Error().Err(err).Msg("bcrypt failure")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to generate password hash",
		})
	}

	user := models.User{
		Username: user_info.Username,
		Password: password_hash,
	}

	if err := query.User.Create(&user); err != nil {
		// Check if user exist. THE BEST ORM I'VE EVER SEEN.
		if sqlite_err, ok := err.(sqlite3.Error); ok {

			switch sqlite_err.Code {
			case sqlite3.ErrConstraint:
				c.JSON(http.StatusConflict, gin.H{
					"message": "username is occupied",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"message":     "unhandled sqlite3 error",
					"sqlite3_err": sqlite_err,
				})
			}

			return
		}

		switch err {
		case gorm.ErrDuplicatedKey:
			c.JSON(http.StatusConflict, gin.H{
				"message": "username is occupied",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
		}
		return
	}

	c.JSON(http.StatusCreated, user)

}
