package router

import (
	"dys-go-starter-project/infrastructures/middlewares"
	v1 "dys-go-starter-project/modules/user/api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUser(r *gin.RouterGroup) {
	api := r.Group("/user", middlewares.BearerAuthMiddleware)
	{
		api.GET("", new(v1.UserController).GetAllUser)
		api.PUT("/", new(v1.UserController).UpdateUser)
		api.DELETE("/", new(v1.UserController).DeleteUser)
		api.PUT("/role", new(v1.UserController).UpdateRole)
		api.GET("/byemail", new(v1.UserController).GetUserByEmail)
	}
}
