package database

import (
	"gorm.io/gorm"
)

// MEMO : for dip

type SqlHandler interface {
	Create(object interface{}) *gorm.DB
	FindAll(object interface{})
	FindByParams(object interface{}, column string, params interface{}) *gorm.DB
	DeleteById(object interface{}, id string)
}
