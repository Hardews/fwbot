/**
 * @Author: Hardews
 * @Date: 2022/10/30 1:17
**/

package router

import (
	"flag"
	"fmt"
	"fwbot/router/controller"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

const ProjectName = "fwbot"

func InitRouter() {
	engine := gin.Default()

	engine.GET("/", controller.Connection)

	err := engine.Run(":8077")
	if err != nil {
		log.Fatalln("run router failed,err:", err)
	}
}

func InitLog() {
	logFileName := flag.String("log", "./fwbot.log", "Log file name")
	flag.Parse()

	// 设置存储的路径
	logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if logErr != nil {
		fmt.Println("Fail to find", *logFile, "disk start Failed")
		os.Exit(1)
	}

	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetPrefix(ProjectName + ":")
}
