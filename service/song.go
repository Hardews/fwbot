/**
 * @Author: Hardews
 * @Date: 2022/11/3 11:32
**/

package service

import (
	"encoding/json"
	"errors"
	"fwbot/config"
	"io"
	"net/http"
	"strconv"
	"strings"

	"fwbot/model"
	"fwbot/tool"
)

const (
	Song = "点歌"
)

var SongUrl = config.Config.Url

// GetSong 处理点歌的函数
func GetSong(msg model.Message) error {
	// 点歌的基本流程是，先处理用户的输入，然后从部署在服务器的一个项目来获取歌曲的id，然后调用cq码来进行发送
	// 一般来说如果出错了可能是服务器的项目毙掉了，然后我简单做了个处理。
	// 这里讲一下，点歌的基本模板应该是 点歌 歌手信息 歌名（歌手信息可选）
	// 给别人点歌的基本模板是 给xx点歌 歌手信息 歌名
	// 这里通过判断前缀，如果前缀没有点歌，可能是给别人点歌或者说是不能理解的关键词，就统一调用给别人点歌来处理这部分。
	if config.Config.Url == "" {
		return DefaultDealFunc(msg)
	}
	if !strings.HasPrefix(msg.Messages, Song) || len(msg.Messages) <= len(Song+" ") {
		return SongTo(msg)
	}

	return parseSong(msg.Messages[7:], tool.Int64ToString(msg.UserId))
}

// SongTo 处理给别人点歌
func SongTo(msg model.Message) error {
	originalStr := strings.Split(msg.Messages, " ")
	if (!strings.HasPrefix(msg.Messages, "给") && !strings.HasSuffix(msg.Messages, "点歌")) || len(originalStr) != 3 {
		return DefaultDealFunc(msg)
	}

	username := strings.TrimSuffix(strings.TrimPrefix(originalStr[0], "给"), "点歌")

	WsPrivateMsg("发送成功", tool.Int64ToString(msg.UserId))
	return parseSong(originalStr[1]+originalStr[2], username)
}

func parseSong(songName, userId string) error {
	var (
		resp   *http.Response
		client = &http.Client{}
	)

	url := SongUrl + "/cloudsearch"
	query := [][2]string{
		{"keywords", songName},
		{"limit", "1"},
	}

	resp, err := client.Do(tool.Get(url, query))
	if err != nil {
		// 这里可能是服务器毙掉了，做一些简单的处理
		WsPrivateMsg("点歌好像出现了一点错误捏，给你点一首我最爱的歌叭", userId)
		WsPrivateMsg("[CQ:music,type=163,id=7214]", userId)
		err = errors.New("do req failed,err:" + err.Error())
		return err
	}

	res, _ := io.ReadAll(resp.Body)

	type T struct {
		Result struct {
			Songs []struct {
				Id int `json:"id"`
			} `json:"songs"`
		} `json:"result"`
	}

	var (
		musicId  string
		jsonData = T{}
	)

	err = json.Unmarshal(res, &jsonData)

	musicId = strconv.Itoa(jsonData.Result.Songs[0].Id)

	WsPrivateMsg("[CQ:music,type=163,id="+musicId+"]", userId)
	return nil
}
