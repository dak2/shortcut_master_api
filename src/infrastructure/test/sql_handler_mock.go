package test

import (
	"gorm.io/gorm"
)

type SqlHandlerMock struct {
	MockCreate          func(obj interface{}) *gorm.DB
	MockFindAll         func(obj interface{})
	MockFindByParams    func(obj interface{}, column string, params interface{}) *gorm.DB
	MockFindAllByParams func(obj interface{}, column interface{}, params interface{}) *gorm.DB
	MockFindAllByParamsWithRelation func(obj interface{}, params []map[string]interface{}, relations []map[string]interface{}) *gorm.DB
	MockDeleteById      func(obj interface{}, id string)
}

func (m *SqlHandlerMock) Create(obj interface{}) *gorm.DB {
	return m.MockCreate(obj)
}

func (m *SqlHandlerMock) FindAll(obj interface{}) {
	m.MockFindAll(obj)
}

func (m *SqlHandlerMock) FindByParams(obj interface{}, column string, params interface{}) *gorm.DB {
	return m.MockFindByParams(obj, column, params)
}

func (m *SqlHandlerMock) FindAllByParams(obj interface{}, column interface{}, params interface{}) *gorm.DB {
	return m.MockFindAllByParams(obj, column, params)
}

func (m *SqlHandlerMock) FindAllByParamsWithRelation(obj interface{}, params []map[string]interface{}, relations []map[string]interface{}) *gorm.DB {
	return m.MockFindAllByParamsWithRelation(obj, params, relations)
}

func (m *SqlHandlerMock) DeleteById(obj interface{}, id string) {
	m.MockDeleteById(obj, id)
}
