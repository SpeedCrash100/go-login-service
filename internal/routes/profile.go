package routes

import (
	"net/http"

	"github.com/SpeedCrash100/go-login-service/internal/database/query"
	my_jwt "github.com/SpeedCrash100/go-login-service/internal/jwt"
	"github.com/SpeedCrash100/go-login-service/pkg/models"

	"github.com/gin-gonic/gin"
)

// Profile returns user's profile if it log in
func Profile(c *gin.Context) {
	user_val, ok := c.Get(my_jwt.IDENTITY_KEY)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get user from jwt token",
		})
		return
	}
	user, _ := user_val.(*models.UserClaims)

	user_db, err := query.User.GetByUsername(user.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "provided username have not found in database",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       user_db.ID,
		"username": user_db.Username,
	})

}
