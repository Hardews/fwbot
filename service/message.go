/**
 * @Author: Hardews
 * @Date: 2022/11/3 11:28
 * @Description:处理消息类型为message的包，目前业务主要集中在message处理
**/

package service

import (
	"fwbot/model"
	"fwbot/tool"
	"math/rand"
	"strings"
)

const AddStr = "添加回复词"

var (
	KeywordStr  = []string{Song, Weather, AddFace, Corn, GetTask, DelTask, AddStr} // 存储相关关键词的切片
	DealFuncMap = map[string]func(msg model.Message) error{                        // 存储相关词对应方法的map
		Song:    GetSong,
		Weather: GetWeather,
		AddFace: AddFaceFunc,
		Corn:    SetCorn,
		GetTask: ShowTasks,
		DelTask: DelTaskFunc,
		AddStr:  AddDefaultReturnStr,
	}
)

// DealWithMsg 处理发来信息的函数
func DealWithMsg(msg model.Message) error {
	// 这里是先遍历关键词切片，如果有就从map中获取到相应的方法
	for _, s := range KeywordStr {
		if strings.Contains(msg.Messages, s) {
			// 这样处理是防止一定的错误
			dealFunc, ok := DealFuncMap[s]
			if !ok {
				return DefaultSelectFunc(msg)
			}
			return dealFunc(msg)
		}
	}

	return DefaultSelectFunc(msg)
}

// DefaultReturn 这里是一些默认的返回文字
var DefaultReturn = []any{
	"行吧行",
	"ok",
}

// DefaultSelectFunc 默认回复函数，不想改名所以用这个
func DefaultSelectFunc(msg model.Message) error {
	WsPrivateMsg(DefaultReturn[rand.Intn(len(DefaultReturn))], tool.Int64ToString(msg.UserId))
	return nil
}

func AddDefaultReturnStr(msg model.Message) error {
	if strings.HasPrefix(msg.Messages, AddStr) {
		return DefaultSelectFunc(msg)
	}

	var res []string
	if res = strings.Split(msg.Messages, " "); len(res) != 2 {
		return DefaultSelectFunc(msg)
	}

	DefaultReturn = append(DefaultReturn, res[1])
	WsPrivateMsg("添加成功", tool.Int64ToString(msg.UserId))
	return nil
}

/*
不当懒狗 决定改回来

// DefaultReturnFunc 默认返回函数，其实可以变成一个,但是但是我懒得改了现在
var (
	DefaultReturnFunc = []func(msg model.Message) error{DefaultStrFunc, DefaultFaceFunc}
)

func DefaultSelectFunc(msg model.Message) error {
	// 这里是随机选择，返回表情包或文字
	c := make(chan int)
	go func() {
		select {
		case c <- 0:
		case c <- 1:
		}
	}()

	return DefaultReturnFunc[<-c](msg)
}

func DefaultStrFunc(msg model.Message) error {
	WsPrivateMsg(DefaultReturn[tool.RandNum(len(DefaultReturn))], tool.Int64ToString(msg.UserId))
	return nil
}

func DefaultFaceFunc(msg model.Message) error {
	if len(FaceStr) == 0 {
		return DefaultStrFunc(msg)
	}
	return SendFace(tool.Int64ToString(msg.UserId), FaceStr[tool.RandNum(len(FaceStr))])
}
*/
