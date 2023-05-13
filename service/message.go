/**
 * @Author: Hardews
 * @Date: 2022/11/3 11:28
 * @Description:处理消息类型为message的包，目前业务主要集中在message处理
**/

package service

import (
	"fwbot/model"
	"strings"
)

var (
	KeywordStr  = []string{Song, Weather}                   // 存储相关关键词的切片
	DealFuncMap = map[string]func(msg model.Message) error{ // 存储相关词对应方法的map
		Song:    GetSong,
		Weather: GetWeather,
	}
)

// DealWithMsg 处理发来信息的函数
func DealWithMsg(msg model.Message) error {
	// 这里是先遍历关键词切片，如果有就从map中获取到相应的方法
	for _, s := range KeywordStr {
		if strings.Contains(msg.Messages, s) {
			dealFunc, ok := DealFuncMap[s]
			if !ok {
				// 没有？调用 gpt
				return DefaultDealFunc(msg)
			}
			return dealFunc(msg)
		}
	}

	return DefaultDealFunc(msg)
}

// DefaultDealFunc 默认回复函数,调用 gpt
func DefaultDealFunc(msg model.Message) error {
	// TODO GPT 调用
	return nil
}
