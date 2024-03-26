package repositories

import (
	"dys-go-starter-project/modules/auth/model"

	"xorm.io/xorm"
)

type AuthUserRepository interface {
	FindByEmail(email string) (*model.AuthUserModel, error)
}

type AuthUserRepositoryImpl struct {
	db *xorm.Engine
}

// Get the user by Email request from API
func (r AuthUserRepositoryImpl) FindByEmail(email string) (*model.AuthUserModel, error) {
	result := model.AuthUserModel{}
	_, err := r.db.Table(model.USER_TABLE_NAME).Where("email = ?", email).Get(&result)
	return &result, err
}
