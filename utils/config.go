package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Mqtt struct {
		Server   string `json:"server"`
		ClientID string `json:"client_id"`
		Sub      string `json:"sub"`
		Pub      string `json:"pub"`
		SysSub   string `json:"sys_sub"`
	} `json:"mqtt"`
}

var ConfigObj = new(Config)

func init() {
	fmt.Println("Init Config")
	readConfigFile()
}

func readConfigFile() (err error) {
	cfgByte, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println("配置文件打开错误")
		return
	}
	//m := make(map[string]interface{})
	err = json.Unmarshal(cfgByte, &ConfigObj)
	return
}
