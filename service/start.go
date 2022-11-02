/**
 * @Author: Hardews
 * @Date: 2022/10/30 0:38
**/

package service

import (
	"fwbot/model"
	"github.com/gorilla/websocket"
	"github.com/robfig/cron/v3"
)

var Conn *websocket.Conn

func Start(conn *websocket.Conn) {
	cr = cron.New(cron.WithSeconds()) //withSeconds精确到秒
	cr.Start()

	Conn = conn

	RChan = make(chan []byte)
	WChan = make(chan model.Action)

	go XianToVCDefaultFunc()
	go Reader()
	go Writer()

	for {
		select {
		case v := <-RChan:
			Classify(v)
		}
	}
}
