package service

import (
	"fwbot/model"
	"fwbot/util"
	"github.com/robfig/cron/v3"
	"log"
	"strings"
)

var (
	cr          *cron.Cron
	defaultTask = map[string]func(){
		"* * 6 * * *":  func() { WsPrivateMsg("陈羡先生提醒您记得晨读和写德语练习册!", util.VcUserId) },
		"* * 18 * * *": func() { WsPrivateMsg("陈羡先生提醒您记得背单词嗷", util.VcUserId) },
	}
)

func init() {
	cr = cron.New(cron.WithSeconds()) //withSeconds精确到秒
}

func XianSetCorn(msg model.Message) error {
	if !strings.HasPrefix(msg.Messages, XianCorn) {
		return DefaultSelectFunc(msg)
	}

	// 这里存在一个问题其实就是不应该由用户输入标准的指令，很容易导致崩溃啊或者很多不安全的问题
	// 但是内部人使用就没那么多说法
	// 先用我的默认处理函数把
	res := strings.Split(msg.Messages, " ")
	if len(res) != 8 {
		WsPrivateMsg("输入有误", util.Int64ToString(msg.UserId))
		return DefaultSelectFunc(msg)
	}
	task := res[1]
	resp := task + "完成了吗"

	specArr := res[2:]

	var spec string
	for _, s := range specArr {
		spec += s
		spec += " "
	}
	spec = strings.TrimSuffix(spec, " ")

	_, err := cr.AddFunc(spec, func() {
		WsPrivateMsg(resp, util.Int64ToString(msg.UserId))
	})
	if err != nil {
		return err
	}
	cr.Start()

	//defer c.Stop() 关了还怎么执行
	return nil
}

// XianToVCDefaultFunc 定时任务处理函数(只需要改上面的map就可以实现改变初始任务
func XianToVCDefaultFunc() {
	for spec, task := range defaultTask {
		_, err := cr.AddFunc(spec, task)
		if err != nil {
			log.Println("begin task failed,err:", err)
			WsPrivateMsg("begin task failed,err:"+err.Error(), util.LenUserId)
			continue
		}
	}
}
