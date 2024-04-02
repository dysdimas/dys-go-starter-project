package services

import (
	inf "dys-go-starter-project/infrastructures"
	"dys-go-starter-project/modules/user/model"
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

// Getting user from repo by email
func (s *UserService) GetUserByEmail(email string) (*model.UserModel, error) {
	userRepository, err := inf.Get[repositories.UserRepository](s.db)
	if err != nil {
		return nil, err
	}

	result, err := userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, errors.New("user not found")
	}
	return result, nil
}

// Update data by email
func (s *UserService) UpdateUser(data *model.UserModel) error {
	userRepository, err := inf.Get[repositories.UserRepository](s.db)
	if err != nil {
		return err
	}

	result, err := userRepository.GetUserByEmail(data.Email)
	if err != nil {
		return err
	}

	if result == nil {
		return errors.New("user not found")
	}

	err = userRepository.UpdateUser(data)
	if err != nil {
		return err
	}

	return nil
}

// Delete data user by email
func (s *UserService) DeleteUser(email string) error {
	userRepository, err := inf.Get[repositories.UserRepository](s.db)
	if err != nil {
		return err
	}

	err = userRepository.DeleteUser(email)
	if err != nil {
		return err
	}

	return nil

}

// Update role user by email
func (s *UserService) UpdateRole(data *model.UserModel) error {
	userRepository, err := inf.Get[repositories.UserRepository](s.db)
	if err != nil {
		return err
	}

	result, err := userRepository.GetUserByEmail(data.Email)
	if err != nil {
		return err
	}

	if result == nil {
		return errors.New("user not found")
	}

	err = userRepository.UpdateRole(data)
	if err != nil {
		return err
	}

	return nil
}
