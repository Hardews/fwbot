/**
 * @Author: Hardews
 * @Date: 2022/10/30 1:37
**/

// 比如hardews.go就是处理我发送的信息的文件,
// 防止冲突，每个函数前应带文件名前三个字符

package service

import (
	"strings"

	"fwbot/model"
)

const (
	HarSongUrl = "http://49.235.99.195:3000"
)

var (
	HarKeywordStr  = []string{}
	HarDealFuncMap = map[string]func(msg model.Message) error{}
)

// HarDealWithMsg 处理lmh发来信息的函数
func HarDealWithMsg(msg model.Message) error {
	// 打个样
	for _, s := range HarKeywordStr {
		if strings.Contains(msg.Messages, s) {
			dealFunc, ok := HarDealFuncMap[s]
			if !ok {
				return DefaultRespFunc(msg)
			}
			return dealFunc(msg)
		}
	}

	return DealWithGeneralMsg(msg)
}
