package controllers

import (
	inf "dys-go-starter-project/infrastructures"
	"dys-go-starter-project/modules/user/api/dtos/impartial"
	"dys-go-starter-project/modules/user/api/dtos/request"
	"dys-go-starter-project/modules/user/api/dtos/response"
	"dys-go-starter-project/modules/user/model"
	"dys-go-starter-project/modules/user/services"
	"dys-go-starter-project/utils/formatter"
	"net/http"

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

func (c *UserController) GetUserByEmail(ctx *gin.Context) {
	email := ctx.Query("email")
	userService, err := inf.Get[*services.UserService](ctx)
	if err != nil {
		inf.Err500ISE(ctx, err.Error())
		return
	}

	user, err := userService.GetUserByEmail(email)
	if err != nil {
		inf.Err404NF(ctx)
		return
	}

	inf.Ok(
		ctx,
		nil,
		&impartial.SuccessImpartial{
			Code:    http.StatusOK,
			Message: "get user successfully",
		},
		&response.UserResponse{
			Id:        formatter.EncryptMd5(string(rune(user.Id))),
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		},
	)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	requestBody := &request.UserUpdateRequest{}
	err := ctx.ShouldBindJSON(requestBody)
	if err != nil {
		inf.Err400BR(ctx, err.Error())
		return
	}

	userService, err := inf.Get[*services.UserService](ctx)
	if err != nil {
		inf.Err500ISE(ctx, err.Error())
	}

	cnvUserModel := &model.UserModel{
		Name:  requestBody.Name,
		Email: requestBody.Email,
	}

	err = userService.UpdateUser(cnvUserModel)
	if err != nil {
		inf.Err404NF(ctx)
		return
	}

	inf.Ok(
		ctx,
		nil,
		&impartial.SuccessImpartial{
			Code:    http.StatusOK,
			Message: "update user successfully",
		},
		nil,
	)

}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	email := ctx.Query("email")
	userService, err := inf.Get[*services.UserService](ctx)
	if err != nil {
		inf.Err500ISE(ctx, err.Error())
		return
	}

	user, err := userService.GetUserByEmail(email)
	if err != nil {
		inf.Err404NF(ctx)
		return
	}

	err = userService.DeleteUser(email)
	if err != nil {
		inf.Err404NF(ctx)
		return
	}

	inf.Ok(
		ctx,
		nil,
		&impartial.SuccessImpartial{
			Code:    http.StatusOK,
			Message: "delete user successfully",
		},
		&response.UserResponse{
			Id:        formatter.EncryptMd5(string(rune(user.Id))),
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		},
	)
}

func (c *UserController) UpdateRole(ctx *gin.Context) {
	requestBody := &request.UserUpdateRequest{}
	err := ctx.ShouldBindJSON(requestBody)
	if err != nil {
		inf.Err400BR(ctx, err.Error())
		return
	}
	userService, err := inf.Get[*services.UserService](ctx)
	if err != nil {
		inf.Err500ISE(ctx, err.Error())
		return
	}

	cnvUserModel := &model.UserModel{
		Role:  requestBody.Role,
		Email: requestBody.Email,
	}
	err = userService.UpdateRole(cnvUserModel)
	if err != nil {
		inf.Err400BR(ctx, err.Error())
		return
	}

	inf.Ok(
		ctx,
		nil,
		&impartial.SuccessImpartial{
			Code:    http.StatusOK,
			Message: "update user role successfully",
		},
		nil,
	)
}
