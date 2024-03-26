package controllers

import (
	inf "dys-go-starter-project/infrastructures"
	"dys-go-starter-project/modules/user/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (c *UserController) GetAllUser(ctx *gin.Context) {
	userService, err := inf.Get[*services.UserService](ctx)
	if err != nil {
		inf.Err500ISE(ctx, err.Error())
		return
	}

	users, err := userService.GetAllUser()
	if err != nil {
		inf.Err500ISE(ctx, err.Error())
		return
	}

	inf.Ok(
		ctx,
		nil,
		nil,
		gin.H{"users": users},
	)
}
