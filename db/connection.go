package db

import (
	"url-shortcut/structs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB = Connect()

func Connect() *gorm.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:zivLevi304974645@tcp(127.0.0.1:3306)/url-shotcut?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	return db
}

func Init() {
	Db.AutoMigrate(&structs.Url{})
}
