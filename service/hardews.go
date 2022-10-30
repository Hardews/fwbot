/**
 * @Author: Hardews
 * @Date: 2022/10/30 1:37
**/

// 比如hardews.go就是处理我发送的信息的文件,
// 防止冲突，每个函数前应带文件名前三个字符

package service

import (
	"encoding/json"
	"fwbot/model"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var HarDefaultReturn = []string{
	"真的吗真的吗",
	"行吧",
	"你说啥是啥",
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
	if !strings.HasPrefix(msg.Messages, HarSong) {
		return HarDefaultFunc()
	}

	var url = HarSongUrl + "/cloudsearch"

	var req *http.Request
	var resp *http.Response
	var client = &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	q := req.URL.Query()
	q.Add("keywords", msg.Messages[7:])
	q.Add("limit", "1")
	req.URL.RawQuery = q.Encode()

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Connection", "keep-alive")

	resp, err = client.Do(req)
	if err != nil {
		log.Println("do req failed,err:", err)
		return err
	}

	defer resp.Body.Close()

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

	return nil
}
