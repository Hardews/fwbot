package service

import (
	"fmt"
	"fwbot/model"
	"fwbot/util"
	"github.com/robfig/cron/v3"
	"log"
	"strconv"
	"strings"
)

var (
	cr          *cron.Cron
	taskMap     = make(map[cron.EntryID]string)
	defaultTask = map[string]func(){
		"0 0 6 * * *":  func() { WsPrivateMsg("陈羡先生提醒您记得晨读和写德语练习册!", util.VcUserId) },
		"0 0 18 * * *": func() { WsPrivateMsg("陈羡先生提醒您记得背单词嗷", util.VcUserId) },
	}
)

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

	taskId, err := cr.AddFunc(spec, func() {
		WsPrivateMsg(resp, util.Int64ToString(msg.UserId))
	})
	if err != nil {
		return err
	}

	taskMap[taskId] = task
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

func XianShowTasks(msg model.Message) error {
	if msg.Messages != XianGetTask {
		return DefaultSelectFunc(msg)
	}

	var res string
	for id, task := range taskMap {
		res += fmt.Sprintf("task id:%d,task name:%s\n", id, task)
	}
	res += "删除任务请通过task id删除"
	WsPrivateMsg(res, util.Int64ToString(msg.UserId))
	return nil
}

func XianDelTaskFunc(msg model.Message) error {
	if !strings.HasPrefix(msg.Messages, XianDelTask) {
		return DefaultSelectFunc(msg)
	}

	taskId, err := strconv.Atoi(msg.Messages[len(XianDelTask):])
	if err != nil {
		return DefaultSelectFunc(msg)
	}

	_, ok := taskMap[cron.EntryID(taskId)]
	if !ok {
		WsPrivateMsg("任务不存在", util.Int64ToString(msg.UserId))
		return nil
	}

	cr.Remove(cron.EntryID(taskId))
	delete(taskMap, cron.EntryID(taskId))
	return nil
}
