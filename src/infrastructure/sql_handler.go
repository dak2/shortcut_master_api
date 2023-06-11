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

func genDbSetting() string {
	conf := config.Init()
	// ref : https://gorm.io/ja_JP/docs/connecting_to_the_database.html#MySQL
	dsn := fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local", conf.DbUser, conf.DbPassword, conf.DbHost, conf.DbName)
	return dsn
}

func NewSqlHandler() *SqlHandler {
	dsn := genDbSetting()
	mySql, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(dsn + "database can't connect")
		panic(err.Error)
	}
	mySql.Set("gorm:table_options", "ENGINE=InnoDB")
	sqlHandler := new(SqlHandler)
	sqlHandler.db = mySql
	return sqlHandler
}

func (handler *SqlHandler) Create(obj interface{}) {
	handler.db.Create(obj)
}

func (handler *SqlHandler) FindAll(obj interface{}) {
	handler.db.Find(obj)
}

func (handler *SqlHandler) DeleteById(obj interface{}, id string) {
	handler.db.Delete(obj, id)
}
