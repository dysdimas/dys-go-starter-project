package middlewares

import (
	"dys-go-starter-project/infrastructures"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("cors middleware recovered", r)
			infrastructures.Err500ISE(c, fmt.Sprintf("%v", r))
		}
	}()

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, app-token, token, client-id, client-version")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
	c.Next()
}
