package services

import (
	inf "dys-go-starter-project/infrastructures"
	"dys-go-starter-project/modules/auth/model"
	"dys-go-starter-project/modules/auth/repositories"
	"dys-go-starter-project/utils/formatter"
	"errors"

	"xorm.io/xorm"
)

type AuthUserService struct {
	db *xorm.Engine
}

func NewAuthService(db *xorm.Engine) *AuthUserService {
	return &AuthUserService{
		db: db,
	}
}

// Getting user by Email
func (s *AuthUserService) GetUserByEmail(email string) (*model.AuthUserModel, error) {
	authUserRepository, err := inf.Get[repositories.AuthUserRepository](s.db)
	if err != nil {
		return nil, err
	}
	result, err := authUserRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, errors.New("user not found")
	}

	return result, nil
}

// Store user
func (s *AuthUserService) SaveUser(data *model.AuthUserModel) (*model.AuthUserModel, error) {
	authUserRepository, err := inf.Get[repositories.AuthUserRepository](s.db)
	if err != nil {
		return nil, err
	}
	result, err := authUserRepository.SaveUser(data)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Validate user
func (s *AuthUserService) ValidateLogin(email string, password string) (*model.AuthUserModel, error) {
	authUserRepository, err := inf.Get[repositories.AuthUserRepository](s.db)
	if err != nil {
		return nil, err
	}

	result, err := authUserRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, errors.New("user not found")
	}

	if result.Password != formatter.EncryptMd5(password) {
		return nil, errors.New("email or password does not match")
	}

	// if result.Password != password {
	// 	return nil, errors.New("email or password does not match")
	// }

	return result, nil
}
