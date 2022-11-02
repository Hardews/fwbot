/**
 * @Author: Hardews
 * @Date: 2022/10/30 0:32
**/

package main

import (
	"fwbot/bot"
	"fwbot/dao"
	"fwbot/router"
)

func main() {
	go bot.Start()
	router.InitLog()
	dao.InitDB()
	router.InitRouter()
}

//func t() {
//	for i := 2231; i < 3000; i++ {
//		f, err := excelize.OpenFile("./weather.xlsx")
//		if err != nil {
//			err = errors.New("open book failed" + err.Error())
//			log.Println(err)
//		}
//
//		defer f.Close()
//
//		// 获取工作表中指定单元格的值
//		code, err := f.GetCellValue("weather", "A"+strconv.Itoa(i))
//		if err != nil {
//			err = errors.New("get book value failed" + err.Error())
//			log.Println(err)
//		}
//
//		name, err := f.GetCellValue("weather", "B"+strconv.Itoa(i))
//		if err != nil {
//			err = errors.New("get book value failed" + err.Error())
//			log.Println(err)
//		}
//		son, err := f.GetCellValue("weather", "C"+strconv.Itoa(i))
//		if err != nil {
//			err = errors.New("get book value failed" + err.Error())
//			log.Println(err)
//		}
//
//		dao.InsertWeather(model.Weather{
//			Code:    code,
//			Name:    name,
//			SonName: son,
//		})
//	}
//	wg.Done()
//}
