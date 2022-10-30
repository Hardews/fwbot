/**
 * @Author: Hardews
 * @Date: 2022/10/30 1:16
**/

package service

import (
	"encoding/json"
	"fmt"
	"fwbot/model"
	"github.com/gorilla/websocket"
	"log"
)

var (
	RChan chan []byte
	WChan chan model.Action
)

func Reader() {
	for true {
		_, msg, err := Conn.ReadMessage()
		if err != nil {
			log.Println("read err:", err)
			return
		}
		RChan <- msg
	}
}

func Writer() {
	for true {
		select {
		case sendData := <-WChan:
			res, err := json.Marshal(&sendData)
			err = Conn.WriteMessage(websocket.TextMessage, res)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
