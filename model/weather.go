/**
 * @Author: Hardews
 * @Date: 2022/10/30 16:49
**/

package model

type Weather struct {
	Uid     int    `gorm:"column:uid"`
	Code    string `gorm:"column:code"`
	Name    string `gorm:"column:name"`
	SonName string `gorm:"column:son"`
}

type WeatherResp struct {
	Status int `json:"status"`
	Result struct {
		Location struct {
			Country  string `json:"country"`
			Province string `json:"province"`
			City     string `json:"city"`
			Name     string `json:"name"`
			Id       string `json:"id"`
		} `json:"location"`
		Now struct {
			Text      string `json:"text"`
			Temp      int    `json:"temp"`
			FeelsLike int    `json:"feels_like"`
			Rh        int    `json:"rh"`
			WindClass string `json:"wind_class"`
			WindDir   string `json:"wind_dir"`
			Uptime    string `json:"uptime"`
		} `json:"now"`
	} `json:"result"`
	Message string `json:"message"`
}
