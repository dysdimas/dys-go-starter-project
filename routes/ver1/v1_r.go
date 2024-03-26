package ver1

import (
	authApi "dys-go-starter-project/modules/auth/router"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) {
	api := r.Group("/v1")
	{
		authApi.RegisterAuth(api)
	}
}
