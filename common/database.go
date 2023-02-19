package common

import (
	"fmt"

	"funboy.top/ginessential/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	// driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "ginessential"
	username := "root"
	password := "root"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", username, password, host, port, database, charset)

	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	// db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connet database, err:" + err.Error())
	}
	db.AutoMigrate(&model.User{})
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
