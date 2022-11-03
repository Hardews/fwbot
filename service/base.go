/**
 * @Author: Hardews
 * @Date: 2022/11/3 14:06
 * @Description:回复相关
**/

package service

import (
	"fwbot/model"
	"fwbot/tool"
)

const Help = "help"

func ShowHelp(msg model.Message) error {
	if msg.Messages != Help {
		return DefaultSelectFunc(msg)
	}
	var res = "现有的关键字有:\n"
	for _, s := range KeywordStr {
		res += s + "\n"
	}
	WsPrivateMsg(res, tool.Int64ToString(msg.UserId))
	return nil
}
