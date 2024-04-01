package repositories

import "dys-go-starter-project/modules/auth/model"

type AuthUserRepository interface {
	FindByEmail(email string) (*model.AuthUserModel, error)
	SaveUser(data *model.AuthUserModel) (*model.AuthUserModel, error)
}
