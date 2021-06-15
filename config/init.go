package config

import (
	"encoding/json"
	"fmt"

	"awesomeProject/ctrl"
	"awesomeProject/soldierInfo"
	"awesomeProject/util"
)
func Init(){
	Umap := JsonInit()
	ctrl.Init(Umap)

}
func JsonInit() map[string]*soldierInfo.Info {
	dataJson := ParseConfigFile()
	var Umap map[string]*soldierInfo.Info

	err := json.Unmarshal(dataJson, &Umap)
	if err != nil {
		fmt.Println("unmarshal json file error")
	}

	da := util.JsonMarshal(Umap)  ///将需要转化的数据通过json.Marshal() 转化为[]byte 格式
	util.JsonWrite(da,"data.json")		  ///将[]byte写入到文件中
	return  Umap
}