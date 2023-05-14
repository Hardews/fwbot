/**
 * @Author: Hardews
 * @Date: 2022/10/30 16:45
**/

package dao

import (
	"fwbot/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	dB *gorm.DB
)

func InitDB() {
	username := config.Config.Username
	password := config.Config.Password
	host := config.Config.Host
	dbName := config.Config.DbName

	dsn := username + ":" + password + "@tcp(" + host + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("failed to connect database")
	}

	dB = db
}
