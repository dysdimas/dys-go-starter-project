package data_utils

import (
	"xorm.io/xorm"
)

func FindByExpression[T Model](db *xorm.Engine, filterExp string, args ...interface{}) ([]T, error) {
	var results []T

	model := *new(T)

	err := db.Table(model.GetTableName()).Where(filterExp, args).Find(&results)
	return results, err
}

func FindOneByExpression[T Model](db *xorm.Engine, filterExp string, args ...interface{}) (T, error) {
	model := *new(T)

	_, err := db.Table(model.GetTableName()).Where(filterExp, args).Get(&model)
	return model, err
}
