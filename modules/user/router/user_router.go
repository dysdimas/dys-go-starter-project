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
		api.GET("/byemail", new(v1.UserController).GetUserByEmail)
		api.PUT("/", new(v1.UserController).UpdateUser)
	}
}
