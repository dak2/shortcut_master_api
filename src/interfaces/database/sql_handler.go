package database

import (
	"gorm.io/gorm"
)

// MEMO : for di

type SqlHandler interface {
	Create(object interface{}) *gorm.DB
	FindAll(object interface{})
	FindByParams(object interface{}, column string, params interface{}) *gorm.DB
	FindAllByParams(object interface{}, column interface{}, params interface{}) *gorm.DB
	FindAllByParamsWithRelation(obj interface{}, params []map[string]interface{}, relations []map[string]interface{}) *gorm.DB
	DeleteById(object interface{}, id string)
}
