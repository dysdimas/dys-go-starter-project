package repositories

import (
	"dys-go-starter-project/modules/auth/model"
	"dys-go-starter-project/utils/formatter"
	"time"

	"xorm.io/xorm"
)

type AuthUserRepositoryImpl struct {
	db *xorm.Engine
}

func NewAuthUserRepositoryImpl(db *xorm.Engine) *AuthUserRepositoryImpl {
	return &AuthUserRepositoryImpl{
		db: db,
	}
}

// Get the user by Email request from API
func (r AuthUserRepositoryImpl) FindByEmail(email string) (*model.AuthUserModel, error) {
	result := model.AuthUserModel{}
	_, err := r.db.Table(model.USER_TABLE_NAME).Where("email = ?", email).Get(&result)
	return &result, err
}

// Save user for register from API
func (r AuthUserRepositoryImpl) SaveUser(data *model.AuthUserModel) (*model.AuthUserModel, error) {
	timestamp := time.Now().Unix()
	t := time.Unix(int64(timestamp), 0)
	newTimestamp := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC).Unix()

	data.Password = formatter.EncryptMd5(data.Password)
	data.CreatedAt = newTimestamp

	_, err := r.db.Table(model.USER_TABLE_NAME).Insert(data)

	if err != nil {
		return nil, err
	}

	return data, nil
}
