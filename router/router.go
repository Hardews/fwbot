/**
 * @Author: Hardews
 * @Date: 2022/10/30 1:17
**/

package router

import (
	"fwbot/router/controller"
	"github.com/gin-gonic/gin"
	"log"
)

func InitRouter() {
	engine := gin.Default()

	engine.GET("/", controller.Connection)

	err := engine.Run(":8077")
	if err != nil {
		log.Fatalln("run router failed,err:", err)
	}
}
