/**
 * @Author: Luminescent0
 * @Date: 2022/11/3 11:30
 * @Description:定时任务相关
**/

package service

import (
	"fmt"
	"fwbot/model"
	"fwbot/tool"
	"github.com/robfig/cron/v3"
	"log"
	"strconv"
	"strings"
)

const (
	Corn    = "定时任务"
	GetTask = "清单"
	DelTask = "删除任务"
)

var (
	cr          *cron.Cron
	taskMap     = make(map[cron.EntryID]string) // 存储定时任务的map
	defaultTask = map[string]func(){            // 默认的定时任务
		"0 0 6 * * *":  func() { WsPrivateMsg("陈羡先生提醒您记得晨读和写德语练习册!", tool.VcUserId) },
		"0 0 18 * * *": func() { WsPrivateMsg("陈羡先生提醒您记得背单词嗷", tool.VcUserId) },
	}
)

func SetCorn(msg model.Message) error {
	if !strings.HasPrefix(msg.Messages, Corn) {
		return DefaultSelectFunc(msg)
	}

	// 这里存在一个问题其实就是不应该由用户输入标准的指令，很容易导致崩溃啊或者很多不安全的问题
	// 但是内部人使用就没那么多说法
	// 先用我的默认处理函数把
	res := strings.Split(msg.Messages, " ")
	if len(res) != 8 {
		WsPrivateMsg("输入有误", tool.Int64ToString(msg.UserId))
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
		WsPrivateMsg(resp, tool.Int64ToString(msg.UserId))
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
			WsPrivateMsg("begin task failed,err:"+err.Error(), tool.LenUserId)
			continue
		}
	}
}

func ShowTasks(msg model.Message) error {
	if msg.Messages != GetTask {
		return DefaultSelectFunc(msg)
	}

	var res string
	for id, task := range taskMap {
		res += fmt.Sprintf("task id:%d,task name:%s\n", id, task)
	}
	res += "删除任务请通过task id删除\n提示:有些完成的可能也在清单上"
	WsPrivateMsg(res, tool.Int64ToString(msg.UserId))
	return nil
}

func DelTaskFunc(msg model.Message) error {
	if !strings.HasPrefix(msg.Messages, DelTask) {
		return DefaultSelectFunc(msg)
	}

	taskId, err := strconv.Atoi(msg.Messages[len(DelTask):])
	if err != nil {
		return DefaultSelectFunc(msg)
	}

	_, ok := taskMap[cron.EntryID(taskId)]
	if !ok {
		WsPrivateMsg("任务不存在", tool.Int64ToString(msg.UserId))
		return nil
	}

	cr.Remove(cron.EntryID(taskId))
	delete(taskMap, cron.EntryID(taskId))
	return nil
}
