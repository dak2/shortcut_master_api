package database

import (
	"log"
	"fmt"
	"short_cut_master_api/src/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func genDbSetting() string {
	conf := config.Config
	// ref : https://gorm.io/ja_JP/docs/connecting_to_the_database.html#MySQL
	dsn := fmt.Sprintf("%s:%s%s/%s?charset=utf8&parseTime=True&loc=Local", conf.DbUser, conf.DbPassword, conf.DbHost, conf.DbName)
	return dsn
}

func openDB() *gorm.DB {
	dsn := genDbSetting()
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(dsn + "database can't connect")
		panic(err.Error)
	}
	DB.Set("gorm:table_options", "ENGINE=InnoDB")
	return DB
}
