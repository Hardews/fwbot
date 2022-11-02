/**
 * @Author: Hardews
 * @Date: 2022/10/30 0:38
**/

package service

import (
	"fwbot/model"
	"github.com/gorilla/websocket"
)

var Conn *websocket.Conn

func Start(conn *websocket.Conn) {
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
