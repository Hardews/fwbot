/**
 * @Author: Hardews
 * @Date: 2022/10/30 11:13
**/

package model

/*
以下是websocket发送消息时需要的json格式
{
    "action": "终结点名称, 例如 'send_group_msg'",
    "params": {
        "参数名": "参数值",
        "参数名2": "参数值"
    },
    "echo": "'回声', 如果指定了 echo 字段, 那么响应包也会同时包含一个 echo 字段, 它们会有相同的值"
}
*/

type Action struct {
	Action string `json:"action"`
	Params any    `json:"params"`
	Echo   string `json:"echo"`
}

type Private struct {
	UserId string `json:"user_id"`
	Msg    string `json:"message"`
}
