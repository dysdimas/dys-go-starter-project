package middlewares

import "github.com/gin-gonic/gin"

func Attach(engine *gin.Engine) {
	engine.Use(CorsMiddleware)
}
