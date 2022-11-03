/**
 * @Author: Hardews
 * @Date: 2022/11/3 11:32
**/

package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"fwbot/model"
	"fwbot/tool"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const (
	SongUrl = "http://49.235.99.195:3000"
	Song    = "点歌"
)

// GetSong 处理点歌的函数
func GetSong(msg model.Message) error {
	// 点歌的基本流程是，先处理用户的输入，然后从我部署在服务器的一个项目来获取歌曲的id，然后调用cq码来进行发送
	// 一般来说如果出错了可能是服务器的项目毙掉了，然后我简单做了个处理。
	// 这里讲一下，点歌的基本模板应该是 点歌 歌手信息 歌名（歌手信息可选）
	// 给别人点歌的基本模板是 给xx点歌 歌手信息 歌名
	// 这里通过判断前缀，如果前缀没有点歌，可能是给别人点歌或者说是不能理解的关键词，就同意调用给别人点歌来处理这部分。
	if !strings.HasPrefix(msg.Messages, Song) || len(msg.Messages) <= len(Song+" ") {
		return SongTo(msg)
	}

	return parseSong(msg.Messages[7:], tool.HarUserId, tool.Int64ToString(msg.UserId))
}

// SongTo 处理给别人点歌
func SongTo(msg model.Message) error {
	originalStr := strings.Split(msg.Messages, " ")
	if (!strings.HasPrefix(msg.Messages, "给") && !strings.HasSuffix(msg.Messages, "点歌")) || len(originalStr) != 3 {
		return DefaultSelectFunc(msg)
	}

	username := strings.TrimSuffix(strings.TrimPrefix(originalStr[0], "给"), "点歌")
	fmt.Println(username)
	userId, ok := tool.RespMap[username]
	if !ok {
		var r = []string{
			"你要不试试给其他人点？？",
			"我我我好像帮不了你捏？",
		}
		WsPrivateMsg(r[tool.RandNum(len(r))], tool.Int64ToString(msg.UserId))
		return nil

	}

	WsPrivateMsg("发送成功", tool.Int64ToString(msg.UserId))
	return parseSong(originalStr[1]+originalStr[2], userId, tool.Int64ToString(msg.UserId))
}

func parseSong(songName, userId, sendId string) error {
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
		WsPrivateMsg("点歌好像出现了错误，err"+err.Error(), sendId)
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
