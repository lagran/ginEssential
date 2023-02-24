package common

import (
	"fmt"

	"funboy.top/ginessential/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// driverName := "mysql"
	// host := "localhost"
	// port := "3306"
	// database := "ginessential"
	// username := "root"
	// password := "root"
	// charset := "utf8"
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	// fmt.Printf("db url: %s:%s", host, port)

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", username, password, host, port, database, charset)

	fmt.Printf("args: %s", args)
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
