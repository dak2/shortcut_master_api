package infrastructure

import (
	"gorm.io/gorm"
)

type SqlHandlerMock struct {
	MockCreate  func(obj interface{}) *gorm.DB
	MockFindAll func(obj interface{})
	MockFindByParams func(obj interface{}, column string, params interface{}) *gorm.DB
	MockDeleteById   func(obj interface{}, id string)
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

func (m *SqlHandlerMock) DeleteById(obj interface{}, id string) {
	m.MockDeleteById(obj, id)
}
