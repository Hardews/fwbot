/**
 * @Author: Hardews
 * @Date: 2022/10/30 0:38
**/

package service

import "github.com/gorilla/websocket"

var Conn *websocket.Conn

func Start(conn *websocket.Conn) {
	Conn = conn

	RChan = make(chan []byte)

	go Reader()

	for {
		select {
		case v := <-RChan:
			Classify(v)
		}
	}
}
