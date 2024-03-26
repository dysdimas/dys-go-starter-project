package router

import (
	v1 "dys-go-starter-project/modules/auth/api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAuth(r *gin.RouterGroup) {
	api := r.Group("/auth")
	{
		api.POST("/register", new(v1.AuthUserController).SaveUser)
		api.POST("/login", new(v1.AuthUserController).Login)
		api.POST("/logout", new(v1.AuthUserController).Logout)
	}

}
