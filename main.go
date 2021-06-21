package main

import (
	"awesomeProject/config"
	"awesomeProject/router"
	"github.com/spf13/pflag"
)

func main(){

	var jsonPath = pflag.StringP("config.army.model.json	-path", "p", "", "Input JsonPath")
	var iniPath = pflag.StringP("ini","i","","input iniPath")
	// 把用户传递的命令行参数解析为对应变量的值
	pflag.Parse()


	config.Init(*jsonPath)
	port := config.ReadConfigPort(*iniPath)  ///通过配置文件获取端口号
	r := router.SetUpRoute()
	r.Run(":"+port)
}




