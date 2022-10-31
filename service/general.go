/**
 * @Author: Hardews
 * @Date: 2022/10/31 20:13
**/

package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"fwbot/dao"
	"fwbot/model"
	"fwbot/util"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var DefaultReturn = []any{}

const (
	Song    = "点歌"
	Weather = "天气"
)

var (
	KeywordStr  = []string{Song, Weather}
	DealFuncMap = map[string]func(msg model.Message) error{
		Song:    GetSong,
		Weather: GetWeather,
	}
)

// DealWithGeneralMsg 处理没有私人定制的人发来信息的函数
func DealWithGeneralMsg(msg model.Message) error {
	// 打个样
	for _, s := range KeywordStr {
		if strings.Contains(msg.Messages, s) {
			dealFunc, ok := DealFuncMap[s]
			if !ok {
				return DefaultRespFunc()
			}
			return dealFunc(msg)
		}
	}

	return DefaultRespFunc()
}

func DefaultRespFunc() error {
	WsPrivateMsg(DefaultReturn[rand.Intn(len(DefaultReturn))], util.HarUserId)
	return nil
}

func GetSong(msg model.Message) error {
	if !strings.HasPrefix(msg.Messages, Song) || len(msg.Messages) <= len(Song+" ") {
		return SongTo(msg)
	}

	return parseSong(msg.Messages[7:], util.HarUserId, util.Int64ToString(msg.UserId))
}

func SongTo(msg model.Message) error {
	originalStr := strings.Split(msg.Messages, " ")
	if (!strings.HasPrefix(msg.Messages, "给") && !strings.HasSuffix(msg.Messages, "点歌")) || len(originalStr) != 3 {
		return DefaultRespFunc()
	}

	username := strings.TrimSuffix(strings.TrimPrefix(originalStr[0], "给"), "点歌")
	fmt.Println(username)
	userId, ok := util.RespMap[username]
	if !ok {
		var r = []string{
			"你要不试试给其他人点？？",
			"我我我好像帮不了你捏？",
			"啥啥啥？",
		}
		WsPrivateMsg(r[rand.Intn(len(r))], util.Int64ToString(msg.UserId))
		return nil
	}

	return parseSong(originalStr[1]+originalStr[2], userId, util.Int64ToString(msg.UserId))
}

func parseSong(songName, userId, sendId string) error {
	var (
		resp   *http.Response
		client = &http.Client{}
	)

	url := HarSongUrl + "/cloudsearch"
	query := [][2]string{
		{"keywords", songName},
		{"limit", "1"},
	}

	resp, err := client.Do(util.Get(url, query))
	if err != nil {
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

	HttpPrivateMsg("发送成功", sendId)
	return HttpPrivateMsg("[CQ:music,type=163,id="+musicId+"]", userId)
}

func GetWeather(msg model.Message) error {
	var city string
	if !strings.HasPrefix(msg.Messages, Weather) {
		return DefaultRespFunc()
	}

	if len(msg.Messages) <= 7 {
		city = "重庆"
	} else {
		city = msg.Messages[7:]
	}

	code := dao.GetCityCode(city)

	if code == "" {
		var r = []string{
			"你要不试试范围小一点？",
			"你你你再说一遍？",
			"啥啥啥？",
		}
		return HttpPrivateMsg(r[rand.Intn(len(r))], util.Int64ToString(msg.UserId))
	}

	url := "https://api.map.baidu.com/weather/v1/"

	query := [][2]string{
		{"district_id", code},
		{"data_type", "now"},
		{"ak", "k4jy5w8xx6yfG76LvLhhmfjpIxzEZrlw"},
	}

	var (
		resp   *http.Response
		client = &http.Client{}
	)

	resp, err := client.Do(util.Get(url, query))
	if err != nil {
		err = errors.New("do req failed,err:" + err.Error())
		return err
	}

	res, _ := io.ReadAll(resp.Body)

	var wea = model.WeatherResp{}

	json.Unmarshal(res, &wea)

	response := "省份:" + wea.Result.Location.Province +
		"\n城市:" + wea.Result.Location.City +
		"\n地区:" + wea.Result.Location.Name +
		"\n时间:" + time.Now().Format("2006/01/02") +
		"\n天气:" + wea.Result.Now.Text +
		"\n风向:" + wea.Result.Now.WindDir +
		"\n风力:" + wea.Result.Now.WindClass
	WsPrivateMsg(response, util.Int64ToString(msg.UserId))

	return nil
}
