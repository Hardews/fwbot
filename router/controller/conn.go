/**
 * @Author: Hardews
 * @Date: 2022/10/30 0:55
**/

package controller

import (
	"fwbot/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
)

var (
	upgrade websocket.Upgrader
)

func Connection(ctx *gin.Context) {
	conn, err := upgrade.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Fatalln("connection bot failed,err :", err)
	}

	service.Start(conn)
}
