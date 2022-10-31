/**
 * @Author: Hardews
 * @Date: 2022/10/30 1:07
**/

package service

import (
	"encoding/json"
	"fmt"
	"fwbot/model"
	"log"
)

// Classify 处理数据的函数
func Classify(msg []byte) {
	var com model.Com
	json.Unmarshal(msg, &com)
	//fmt.Println(string(msg))

	switch com.PostType {
	case "meta_event":
		// 元事件 类似心跳 目前设置的是30s
		fmt.Println("pong")
		break
	case "message":
		// 用户发送过来的消息
		// 主要是处理这里的信息，当有人发信息时应该返回什么之类的
		var m model.Message
		json.Unmarshal(msg, &m)

		// 这里打个样
		// 为了保证代码简洁，所以这里采用这种方式，处理消息的方法统一放在deal层
		// 以下代码改动应该只在map元素增添删减。
		dealFunc, ok := map[int64]func(msg model.Message) error{
			1225101127: HarDealWithMsg,
			3332648553: HarDealWithMsg,
		}[m.UserId]
		if !ok {
			err := DealWithGeneralMsg(m)
			if err != nil {
				log.Println("回复出错，err:", err)
			}
			return
		}

		err := dealFunc(m)
		if err != nil {
			log.Println("回复出错，err:", err)
			return
		}
	case "request":
	case "notice":
		var m model.Message
		json.Unmarshal(msg, &m)
		fmt.Println(m)
	}
}
