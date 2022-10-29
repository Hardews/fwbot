/**
 * @Author: Hardews
 * @Date: 2022/10/30 1:46
**/

// 此包封装的是返回信息的函数

package send

import (
	"net/http"
)

const base = "http://127.0.0.1:8077" // 服务器地址

func SendPrivateMsg(msg, userId string) error {
	url := base + "/send_private_msg"

	_, err := http.Get(url + "?user_id=" + userId + "&&message=" + msg)
	return err
}
