package main

import (
	"awesomeProject/config"
	"awesomeProject/router"
)

func main(){

	config.Init()
	port := config.ReadConfigPort("./app.ini")  ///通过配置文件获取端口号
	r := router.SetUpRoute()
	r.Run(":"+port)
}




