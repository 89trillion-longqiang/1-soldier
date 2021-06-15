package util

import (
	"encoding/json"
	"log"
	"os"

	"awesomeProject/soldierInfo"
)

// JsonMarshal /将数据转化为[]byte
func JsonMarshal(info map[string]*soldierInfo.Info) []byte {

	data, err := json.Marshal(info)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

// JsonWrite /将数据写入json文件中，生成新的json文件
func JsonWrite(data []byte,filename string) {
	fp, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	_, err = fp.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}
