package repositories

import "dys-go-starter-project/modules/user/model"

type UserRepository interface {
	GetAllUser() (*[]map[string]interface{}, error)
	GetUserByEmail(email string) (*model.UserModel, error)
	UpdateUser(*model.UserModel) error
	DeleteUser(email string) error
}
