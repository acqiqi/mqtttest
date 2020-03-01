package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func init() {
	log.Println("Init File")
}

func writeFile(fileName string, data interface{}) (err error) {
	b, _ := json.Marshal(data)
	err = ioutil.WriteFile("./config/"+fileName+".json", b, 0666)
	return
}

func ReadFile(fileName string) (data string, err error) {
	dataBytes, err := ioutil.ReadFile("./jsonfiles/" + fileName + ".json")
	if err != nil {
		return "", err
	}
	//m := make(map[string]interface{})
	return string(dataBytes), nil
}
