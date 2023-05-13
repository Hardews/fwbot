/**
 * @Author: Hardews
 * @Date: 2022/10/30 0:38
**/

package service

import (
	"encoding/json"
	"fmt"
	"log"

	"fwbot/model"

	"github.com/gorilla/websocket"
)

var Conn *websocket.Conn

func Start(conn *websocket.Conn) {
	Conn = conn

	RChan = make(chan []byte)
	WChan = make(chan model.Action)

	go Reader()
	go Writer()

	for {
		select {
		case v := <-RChan:
			Classify(v)
		}
	}
}

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
