/**
 * @Author: Hardews
 * @Date: 2022/11/3 11:33
**/

package service

import (
	"encoding/json"
	"errors"
	"fwbot/dao"
	"fwbot/model"
	"fwbot/tool"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	Weather    = "天气"
	WeatherUrl = "https://api.map.baidu.com/weather/v1/"
)

func GetWeather(msg model.Message) error {
	var city string
	if !strings.HasPrefix(msg.Messages, Weather) {
		return DefaultSelectFunc(msg)
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
		}
		WsPrivateMsg(r[tool.RandNum(len(r))], tool.Int64ToString(msg.UserId))
		return nil
	}

	query := [][2]string{
		{"district_id", code},
		{"data_type", "now"},
		{"ak", "k4jy5w8xx6yfG76LvLhhmfjpIxzEZrlw"},
	}

	var (
		resp   *http.Response
		client = &http.Client{}
	)

	resp, err := client.Do(tool.Get(WeatherUrl, query))
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
	WsPrivateMsg(response, tool.Int64ToString(msg.UserId))
	return nil
}
