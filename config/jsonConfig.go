package config

import (
	"fmt"
	"io/ioutil"
)

func ParseConfigFile(filename string) []byte{


	dataJson, err := ioutil.ReadFile(filename) //读取json文件
	if err != nil {
		fmt.Println("read file error")
		return nil
	}
	return dataJson
}