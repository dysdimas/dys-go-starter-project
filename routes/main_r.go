package routes

import (
	"dys-go-starter-project/infrastructures"
	"dys-go-starter-project/infrastructures/middlewares"
	ver1 "dys-go-starter-project/routes/ver1"
	"time"

	"github.com/gin-gonic/gin"
)

func Attach(engine *gin.Engine) {
	engine.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"time": time.Now(),
		})
	})

	apiGroup := engine.Group("/api", middlewares.ApiGroup)
	{
		ver1.RegisterRoutes(apiGroup)
	}

	engine.NoRoute(func(c *gin.Context) {
		infrastructures.Err404NF(c)
	})
}
