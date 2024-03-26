package data_utils

import (
	"dys-go-starter-project/utils"

	"xorm.io/xorm"
)

func Create[T Model](db *xorm.Engine, data T) error {
	m, err := utils.AnyToMap(data)
	if err != nil {
		return err
	}

	delete(m, "id")
	model := *new(T)
	_, err = db.Table(model.GetTableName()).Insert(m)
	return err
}

func UpdateById[T Model](db *xorm.Engine, id int64, m map[string]interface{}) error {
	delete(m, "id")
	model := *new(T)
	_, err := db.Table(model.GetTableName()).ID(id).Update(m)
	return err
}

func DeleteById[T Model](db *xorm.Engine, id int64) error {
	model := *new(T)
	_, err := db.Table(model.GetTableName()).Delete(id)
	return err
}
