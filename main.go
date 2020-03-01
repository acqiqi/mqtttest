package main

import (
	"fmt"
	"mqtttest/utils"
	"os"
	"time"
)

func main() {
	fmt.Println("MQTT Test Init")
	//fmt.Println(utils.ConfigObj)
	arg_num := len(os.Args)
	if arg_num != 2 {
		fmt.Println("Error : Not FilePath")
		return
	}
	fmt.Println(os.Args[1])
	str, err := utils.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(str)
	time.Sleep(time.Second)
	utils.Publish(utils.ConfigObj.Mqtt.Pub, 1, str)
	time.Sleep(3 * time.Second)

}
