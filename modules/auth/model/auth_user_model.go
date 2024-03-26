package model

import (
	"dys-go-starter-project/modules/auth/api/dtos/request"
)

const USER_TABLE_NAME = "users"

type AuthUserModel struct {
	Id        int64  `xorm:"not null pk autoincr 'id'"`
	Name      string `xorm:"not null 'name'"`
	Email     string `xorm:"not null 'email'"`
	Password  string `xorm:"not null 'password'"`
	CreatedAt int64  `xorm:"not null 'created_at'"`
	UpdatedAt *int64 `xorm:"null 'updated_at'"`
}

func (m *AuthUserModel) TableName() string {
	return USER_TABLE_NAME
}

func ConvertToAuthUserModel(requestBody *request.RegisterRequest) *AuthUserModel {
	return &AuthUserModel{
		Name:     requestBody.Name,
		Email:    requestBody.Email,
		Password: requestBody.Password,
	}
}
