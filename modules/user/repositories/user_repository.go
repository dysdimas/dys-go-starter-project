package repositories

import (
	"dys-go-starter-project/modules/user/model"

	"xorm.io/xorm"
)

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
	checkUser, err := r.db.Table(model.USER_TABLE_NAME).Where("email = ?", email).Get(&results)
	if err != nil {
		return nil, err
	}

	if !checkUser {
		return nil, err
	}

	return &results, err
}

// Update data name by email
func (r UserRepositoryImpl) UpdateUser(data *model.UserModel) error {
	_, err := r.db.Table(model.USER_TABLE_NAME).Update(&model.UserModel{
		Name: data.Name,
	}, &model.UserModel{Email: data.Email})
	if err != nil {
		return err
	}
	return nil
}

// Delete user by email
func (r UserRepositoryImpl) DeleteUser(email string) error {
	_, err := r.db.Table(model.USER_TABLE_NAME).Where("email = ?", email).Delete()
	if err != nil {
		return err
	}

	return nil
}
