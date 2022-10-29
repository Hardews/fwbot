/**
 * @Author: Hardews
 * @Date: 2022/10/30 1:16
**/

package service

import (
	"log"
)

var (
	RChan chan []byte
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
