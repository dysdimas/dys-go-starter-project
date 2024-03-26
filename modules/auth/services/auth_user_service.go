package services

import (
	"crypto/md5"
	inf "dys-go-starter-project/infrastructures"
	"dys-go-starter-project/modules/auth/model"
	"dys-go-starter-project/modules/auth/repositories"
	"errors"
	"fmt"

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

	// if result.Password != s.encryptMd5(password) {
	// 	return nil, errors.New("email or passwor does not match")
	// }

	if result.Password != password {
		return nil, errors.New("email or passwor does not match")
	}

	return result, nil
}

func (s *AuthUserService) encryptMd5(word string) string {
	h := md5.New()
	return fmt.Sprintf("%x", h.Sum([]byte(word)))
}
