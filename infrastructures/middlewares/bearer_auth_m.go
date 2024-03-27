package middlewares

import (
	"dys-go-starter-project/infrastructures"
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func BearerAuthMiddleware(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("bearer auth middleware recovered", r)
			infrastructures.Err500ISE(c, fmt.Sprintf("%v", r))
		}
	}()
	authorizationHeader := c.GetHeader(infrastructures.HeaderAuthorization)
	if authorizationHeader == "" {
		c.Abort()
		infrastructures.Err400BR(c, "endpoint restricted")
		c.Writer.Header().Set("WWW-Authenticate", "Bearer auth=Restricted")
		return
	}

	if !strings.Contains(authorizationHeader, infrastructures.AuthMethodBearer) {
		c.Abort()
		infrastructures.Err400BR(c, "endpoint restricted")
		c.Writer.Header().Set("WWW-Authenticate", "Bearer auth=Restricted")
		return
	}

	//audience := c.GetHeader(infrastructures.HeaderAppClient)
	//if audience == "" {
	//	c.Abort()
	//	c.Writer.Header().Set("WWW-Authenticate", "Bearer auth=Restricted")
	//	infrastructures.Err400BR(c, "endpoint restricted")
	//	return
	//}

	//cfgAudiences := os.Getenv(infrastructures.EnvJwtAudiences)
	//if cfgAudiences == "" {
	//	c.Abort()
	//	infrastructures.Err500ISE(c, "audiences does not exist")
	//	return
	//}
	//
	//audiences := strings.Split(cfgAudiences, "|")
	//isContain := utils.SliceContains(audiences, audience)
	//if !isContain {
	//	c.Abort()
	//	c.Writer.Header().Set("WWW-Authenticate", "Bearer auth=Restricted")
	//	infrastructures.Err400BR(c, "endpoint restricted")
	//	return
	//}

	tokenString := strings.Replace(authorizationHeader, fmt.Sprintf("%s ", infrastructures.AuthMethodBearer), "", -1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("Signing method invalid")
		}

		return []byte(os.Getenv(infrastructures.EnvJwtSecret)), nil
	})

	if err != nil {
		infrastructures.Err400BR(c, err.Error())
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		infrastructures.Err401Unauthorized(c, "token is not valid")
		return
	}

	c.Set(infrastructures.CtxClaims, claims)
	//c.Set(infrastructures.CtxAudience, audience)
}
