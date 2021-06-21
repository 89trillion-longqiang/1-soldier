package config

import (
	goini "github.com/clod-moon/goconf"
)

// ReadConfigPort /读取配置文件来获取端口号
func ReadConfigPort(filePath string)(port string){

	conf := goini.InitConfig(filePath)
	port = conf.GetValue("server", "HttpPort")
	return port
}
