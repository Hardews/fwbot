/**
 * @Author: Hardews
 * @Date: 2022/10/30 16:45
**/

package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	dB *gorm.DB
)

func InitDB() {
	dsn := "root:lmh123@tcp(127.0.0.1:3306)/fwbot?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("failed to connect database")
	}

	dB = db
}
