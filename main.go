/**
 * @Author: Hardews
 * @Date: 2022/10/30 0:32
**/

package main

import (
	"fwbot/dao"
	"fwbot/router"
)

func main() {
	router.InitLog()
	dao.InitDB()
	router.InitRouter()
}
