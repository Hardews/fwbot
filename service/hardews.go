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

var (
	FaceStr        = []string{}
	HarKeywordStr  = []string{Song, Weather, HarAddFace, XianCorn, XianGetTask, XianDelTask}
	HarDealFuncMap = map[string]func(msg model.Message) error{
		Song:        GetSong,
		Weather:     GetWeather,
		HarAddFace:  HarAddFaceFunc,
		XianCorn:    XianSetCorn,
		XianGetTask: XianShowTasks,
		XianDelTask: XianDelTaskFunc,
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

	return DefaultSelectFunc(msg)
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
	WsPrivateMsg("添加成功！", util.Int64ToString(msg.UserId))
	return nil
}
