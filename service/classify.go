/**
 * @Author: Hardews
 * @Date: 2022/10/30 1:07
**/

package service

import (
	"encoding/json"
	"fmt"
	"qq-bot/deal"
	"qq-bot/model"
	"qq-bot/send"
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
		switch m.UserId {
		case 1225101127:
			// 比如说是我发过来的信息
			resp := deal.HarDealWithMsg(m)

			err := send.SendPrivateMsg(resp, "1225101127")
			fmt.Println(err)
		}
	case "request":
	case "notice":
		var m model.Message
		json.Unmarshal(msg, &m)
		fmt.Println(m)
	}
}
