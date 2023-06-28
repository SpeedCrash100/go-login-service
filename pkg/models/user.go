package models

type UserLoginInfo struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type UserClaims struct {
	Username string `json:"username"`
}
