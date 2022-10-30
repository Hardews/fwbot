/**
 * @Author: Hardews
 * @Date: 2022/10/30 1:46
**/

// 此包封装的是返回信息的函数

package service

import (
	"fwbot/model"
	"net/http"
)

const base = "http://127.0.0.1:8078" // http服务器地址

// HttpPrivateMsg 通过http发送私聊消息
func HttpPrivateMsg(msg, userId string) error {
	url := base + "/send_private_msg"

	_, err := http.Get(url + "?user_id=" + userId + "&&message=" + msg)
	return err
}

// WsPrivateMsg 通过WS发送私聊消息
func WsPrivateMsg(msg, userId string) {
	WChan <- model.Action{
		Action: "send_private_msg",
		Params: model.Private{
			UserId: userId,
			Msg:    msg,
		},
		Echo: "",
	}
}
