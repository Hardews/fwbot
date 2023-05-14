/**
 * @Author: Hardews
 * @Date: 2022/10/30 0:32
**/

package main

import (
	"fwbot/config"
	"fwbot/dao"
	"fwbot/router"
)

func main() {
	config.SetConfig()
	dao.InitDB()
	router.InitRouter()
}
