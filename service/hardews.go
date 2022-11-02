/**
 * @Author: Hardews
 * @Date: 2022/10/30 1:37
**/

// 比如hardews.go就是处理我发送的信息的文件,
// 防止冲突，每个函数前应带文件名前三个字符

package service

import (
	"fwbot/util"
	"strings"

	"fwbot/model"
)

const (
	HarSongUrl = "http://49.235.99.195:3000"
	HarAddFace = "添加表情包"
)

var (
	FaceStr        = []string{}
	HarKeywordStr  = []string{HarAddFace, XianCorn}
	HarDealFuncMap = map[string]func(msg model.Message) error{
		HarAddFace: HarAddFaceFunc,
		XianCorn:   XianSetCorn,
	}
)

// HarDealWithMsg 处理lmh发来信息的函数
func HarDealWithMsg(msg model.Message) error {
	// 打个样
	for _, s := range HarKeywordStr {
		if strings.Contains(msg.Messages, s) {
			dealFunc, ok := HarDealFuncMap[s]
			if !ok {
				return DefaultSelectFunc(msg)
			}
			return dealFunc(msg)
		}
	}

	return DealWithGeneralMsg(msg)
}

func HarAddFaceFunc(msg model.Message) error {
	if !strings.HasPrefix(msg.Messages, HarAddFace+"[CQ:image,file=") {
		return DefaultSelectFunc(msg)
	}

	b := strings.Index(msg.Messages, "url=")
	if b == -1 {
		return DefaultSelectFunc(msg)
	}

	url := strings.Split(msg.Messages[b+4:len(msg.Messages)-1], "?")[0]
	FaceStr = append(FaceStr, url)
	return HttpPrivateMsg("添加成功！", util.Int64ToString(msg.UserId))
}
