/**
 * @Author: Hardews
 * @Date: 2022/10/30 16:48
**/

package dao

//func InsertWeather(wea model.Weather) {
//	dB.Create(wea)
//}

func GetCityCode(cityName string) string {
	var res string
	dB.Raw("SELECT code FROM weathers WHERE name LIKE ? OR son LIKE ?", cityName, cityName).Scan(&res)
	return res
}
