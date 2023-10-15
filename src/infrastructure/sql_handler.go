package infrastructure

import (
	"fmt"
	"log"
	config "shortcut_master_api/src/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SqlHandler struct {
	db *gorm.DB
}

type SqlHandlerInterface interface {
	Create(obj interface{}) *gorm.DB
	FindAll(obj interface{})
	FindByParams(obj interface{}, column string, params interface{}) *gorm.DB
	FindAllByParams(obj interface{}, column string, params interface{}) *gorm.DB
	DeleteById(obj interface{}, id string)
}

func genDbSetting() string {
	conf := config.Init()
	// ref : https://gorm.io/ja_JP/docs/connecting_to_the_database.html#MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.DBUser, conf.DBPassword, conf.DBHost, conf.DBName)
	return dsn
}

func NewSqlHandler() *SqlHandler {
	dsn := genDbSetting()
	mySql, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(dsn + "database can't connect")
		panic(err)
	}
	mySql.Set("gorm:table_options", "ENGINE=InnoDB")
	sqlHandler := new(SqlHandler)
	sqlHandler.db = mySql
	return sqlHandler
}

func (handler *SqlHandler) Create(obj interface{}) *gorm.DB {
	return handler.db.Create(obj)
}

func (handler *SqlHandler) FindAll(obj interface{}) {
	handler.db.Find(obj)
}

func (handler *SqlHandler) FindByParams(obj interface{}, column string, params interface{}) *gorm.DB {
	columnCondition := fmt.Sprintf("%s = ?", column)
	res := handler.db.First(obj, columnCondition, params)
	return res
}

func (handler *SqlHandler) FindAllByParams(obj interface{}, column string, params interface{}) *gorm.DB {
	columnCondition := fmt.Sprintf("%s = ?", column)
	res := handler.db.Where(columnCondition, params).Find(obj)
	return res
}

func (handler *SqlHandler) DeleteById(obj interface{}, id string) {
	handler.db.Delete(obj, id)
}
