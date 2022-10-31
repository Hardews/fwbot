package service

import (
	"fwbot/model"
	"github.com/robfig/cron/v3"
	"strings"
)

const (
	XianUserId = "2945008714"
	XianCorn   = "定时任务"
)

func XianDefaultFunc() error {
	//rand.Seed(time.Now().Unix())
	WsPrivateMsg("听不懂捏", XianUserId)
	return nil
}
func XianSetCorn(msg model.Message) error {
	if !strings.HasPrefix(msg.Messages, XianCorn) {
		return XianDefaultFunc()
	}
	res := strings.Split(msg.Messages, " ")
	task := res[1]
	spec := res[2]
	resp := task + "完成了吗"
	c := cron.New(cron.WithSeconds()) //withSeconds精确到秒
	_, err := c.AddFunc(spec, func() {
		WsPrivateMsg(resp, XianUserId)
	})
	if err != nil {
		return err
	}
	c.Start()
	defer c.Stop()
	return nil
}
