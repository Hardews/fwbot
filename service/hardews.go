/**
 * @Author: Hardews
 * @Date: 2022/10/30 1:37
**/

// 比如hardews.go就是处理我发送的信息的文件,
// 防止冲突，每个函数前应带文件名前三个字符

package service

import (
	"encoding/json"
	"errors"
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

var HarDefaultReturn = []any{
	"真的吗真的吗",
	"行吧",
	"你说啥是啥",
	model.QQFaceMsg{
		Type: "face",
		Data: struct {
			Id string `json:"id"`
		}{Id: "1"},
	},
	model.QQFaceMsg{
		Type: "face",
		Data: struct {
			Id string `json:"id"`
		}{Id: "110"},
	},
}

const (
	HarUserId  = "1225101127"
	HarSongUrl = "http://49.235.99.195:3000"
	HarSong    = "点歌"
	HarWeather = "天气"
)

var (
	HarKeywordStr  = []string{HarSong, HarWeather}
	HarDealFuncMap = map[string]func(msg model.Message) error{
		HarSong:    HarGetSong,
		HarWeather: HarGetWeather,
	}
)

// HarDealWithMsg 处理lmh发来信息的函数
func HarDealWithMsg(msg model.Message) error {
	// 打个样
	for _, s := range HarKeywordStr {
		if strings.Contains(msg.Messages, s) {
			dealFunc, ok := HarDealFuncMap[s]
			if !ok {
				return HarDefaultFunc()
			}
			return dealFunc(msg)
		}
	}

	return HarDefaultFunc()
}

func HarDefaultFunc() error {
	rand.Seed(time.Now().Unix())
	WsPrivateMsg(HarDefaultReturn[rand.Intn(len(HarDefaultReturn))], HarUserId)
	return nil
}

func HarGetSong(msg model.Message) error {
	if !strings.HasPrefix(msg.Messages, HarSong) || len(msg.Messages) <= 7 {
		return HarDefaultFunc()
	}

	var (
		resp   *http.Response
		client = &http.Client{}
	)

	url := HarSongUrl + "/cloudsearch"
	query := [][2]string{
		{"keywords", msg.Messages[7:]},
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

	return HttpPrivateMsg("[CQ:music,type=163,id="+musicId+"]", HarUserId)
}

func HarGetWeather(msg model.Message) error {
	var city string
	if !strings.HasPrefix(msg.Messages, HarWeather) {
		return HarDefaultFunc()
	}

	if len(msg.Messages) <= 7 {
		city = "重庆"
	} else {
		city = msg.Messages[7:]
	}

	code := dao.GetCityCode(city)

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

	var wea model.WeatherResp

	json.Unmarshal(res, &wea)

	response := "省份:" + wea.Result.Location.Province +
		"\n城市:" + wea.Result.Location.City +
		"\n地区:" + wea.Result.Location.Name +
		"\n时间:" + time.Now().Format("2006/01/02") +
		"\n天气:" + wea.Result.Now.Text +
		"\n风向:" + wea.Result.Now.WindDir +
		"\n风力:" + wea.Result.Now.WindClass
	WsPrivateMsg(response, HarUserId)

	return nil
}
