package main

import (
	"jdudp/Console"
	"jdudp/routers"
	"jdudp/utils"
)

func main() {
	utils.InitRedisUtil("120.55.54.57", 6379, "zhutao@123")
	router := routers.InitRouter()
	// main.go中关闭定时任务
	defer Console.Conrs.Stop()
	//静态资源
	router.Run(":8082")
}
