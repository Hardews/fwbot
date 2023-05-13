/**
 * @Author: Hardews
 * @Date: 2022/11/3 11:38
 * @Description:发送消息相关
**/

package service

import "fwbot/model"

// WsPrivateMsg 通过WS发送私聊消息
func WsPrivateMsg(msg any, userId string) {
	WChan <- model.Action{
		Action: "send_private_msg",
		Params: model.Private{
			UserId: userId,
			Msg:    msg,
		},
		Echo: "",
	}
}

/*
base        = "http://127.0.0.1:8078" // http服务器地址

// HttpPrivateMsg 通过http发送私聊消息
func HttpPrivateMsg(msg, userId string) error {
	url := base + "/send_private_msg"

	_, err := http.Get(url + "?user_id=" + userId + "&&message=" + msg)
	return err
}
*/
