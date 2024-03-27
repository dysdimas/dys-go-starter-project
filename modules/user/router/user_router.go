package router

import (
	v1 "dys-go-starter-project/modules/user/api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUser(r *gin.RouterGroup) {
	api := r.Group("/user")
	{
		api.GET("", new(v1.UserController).GetAllUser)
		api.GET("/byemail", new(v1.UserController).GetUserByEmail)
	}
}
