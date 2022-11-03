/**
 * @Author: Hardews
 * @Date: 2022/11/3 11:31
 * @Description:表情包相关
**/

package service

import (
	"fwbot/model"
	"fwbot/tool"
	"strings"
)

const AddFace = "添加表情包"

// SendFace 发送表情包的封装
func SendFace(userId, url string) error {
	WsPrivateMsg("[CQ:image,file="+url+",type=show,value=1]", userId)
	return nil
}

// AddFaceFunc 添加表情包处理函数
func AddFaceFunc(msg model.Message) error {
	if !strings.HasPrefix(msg.Messages, AddFace+"[CQ:image,file=") {
		return DefaultSelectFunc(msg)
	}

	b := strings.Index(msg.Messages, "url=")
	if b == -1 {
		return DefaultSelectFunc(msg)
	}

	url := strings.Split(msg.Messages[b+4:len(msg.Messages)-1], "?")[0]
	FaceStr = append(FaceStr, url)
	WsPrivateMsg("添加成功！", tool.Int64ToString(msg.UserId))
	return nil
}
