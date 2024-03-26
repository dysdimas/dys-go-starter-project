package controllers

import (
	inf "dys-go-starter-project/infrastructures"
	"dys-go-starter-project/modules/auth/api/dtos/impartial"
	"dys-go-starter-project/modules/auth/api/dtos/request"
	"dys-go-starter-project/modules/auth/api/dtos/response"
	"dys-go-starter-project/modules/auth/services"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"xorm.io/xorm/convert"
)

type AuthUserController struct {
}

func (c *AuthUserController) Login(ctx *gin.Context) {
	requestBody := &request.LoginRequest{}
	err := ctx.ShouldBindJSON(requestBody)
	if err != nil {
		inf.Err400BR(ctx, err.Error())
		return
	}

	authUserService, err := inf.Get[*services.AuthUserService](ctx)
	if err != nil {
		inf.Err500ISE(ctx, err.Error())
	}

	validateLogin, err := authUserService.ValidateLogin(requestBody.Email, requestBody.Password)
	if err != nil {
		inf.Err401Unauthorized(ctx, "invalid user")
		return
	}

	cfgDuration := os.Getenv(inf.EnvJwtTokenExpirationDuration)
	duration := time.Minute * 5
	cfgDurationInt, err := convert.AsInt64(cfgDuration)
	if err != nil {
		duration = time.Duration(cfgDurationInt) * time.Minute
	}
	createdAt := time.Now()
	expiredAt := createdAt.Add(duration)

	claims := impartial.ClaimImpartial{
		StandardClaims: jwt.StandardClaims{
			Issuer:    os.Getenv(inf.EnvAppName),
			IssuedAt:  createdAt.Unix(),
			ExpiresAt: expiredAt.Unix(),
		},
		Name:  validateLogin.Name,
		Email: validateLogin.Email,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	signedToken, err := token.SignedString([]byte(os.Getenv(inf.EnvJwtSecret)))
	if err != nil {
		inf.Err400BR(ctx, err.Error())
		return
	}

	inf.Ok(
		ctx,
		&response.AuthResponse{
			signedToken, createdAt.Unix(),
			expiredAt.Unix(),
		},
		nil,
	)
}

func (c *AuthUserController) Logout(ctx *gin.Context) {
	inf.Ok(ctx, nil, nil)
}
