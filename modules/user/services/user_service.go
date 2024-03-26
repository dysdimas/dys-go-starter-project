package services

import (
	inf "dys-go-starter-project/infrastructures"
	"dys-go-starter-project/modules/user/repositories"
	"errors"

	"xorm.io/xorm"
)

type UserService struct {
	db *xorm.Engine
}

func NewUserService(db *xorm.Engine) *UserService {
	return &UserService{
		db: db,
	}
}

// Getting all user from user repo
func (s *UserService) GetAllUser() (*[]map[string]interface{}, error) {
	userRepository, err := inf.Get[repositories.UserRepository](s.db)
	if err != nil {
		return nil, err
	}

	result, err := userRepository.GetAllUser()
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, errors.New("user not found in the list")
	}

	return result, nil
}
