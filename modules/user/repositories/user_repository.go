package repositories

import (
	"dys-go-starter-project/modules/user/model"

	"xorm.io/xorm"
)

type UserRepository interface {
	GetAllUser() (*[]map[string]interface{}, error)
	GetUserByEmail(email string) (*model.UserModel, error)
}

type UserRepositoryImpl struct {
	db *xorm.Engine
}

func NewUserRepositoryImpl(db *xorm.Engine) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

// Get all user
func (r UserRepositoryImpl) GetAllUser() (*[]map[string]interface{}, error) {
	var results []map[string]interface{}
	err := r.db.Table(model.USER_TABLE_NAME).Find(&results)
	if err != nil {
		return nil, err
	}
	return &results, err
}

// Get by email
func (r UserRepositoryImpl) GetUserByEmail(email string) (*model.UserModel, error) {
	results := model.UserModel{}
	_, err := r.db.Table(model.USER_TABLE_NAME).Where("email = ?", email).Get(&results)
	if err != nil {
		return nil, err
	}
	return &results, err
}
