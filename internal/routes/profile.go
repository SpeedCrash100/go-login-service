package routes

import (
	"net/http"

	my_jwt "github.com/SpeedCrash100/go-login-service/internal/jwt"
	"github.com/SpeedCrash100/go-login-service/pkg/models"

	"github.com/gin-gonic/gin"
)

// Profile returns user's profile if it log in
func Profile(c *gin.Context) {
	user_val, ok := c.Get(my_jwt.IDENTITY_KEY)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Failed to get user from jwt token",
		})
	}
	user, _ := user_val.(*models.UserClaims)

	c.JSON(http.StatusOK, gin.H{
		"username": user.Username,
	})

}
