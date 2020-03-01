package utils

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"log"
	"time"
)

var Mqttclient mqtt.Client
var MqttToken mqtt.Token

var isOpen = false

func init() {
	//mqtt.DEBUG = log.New(os.Stdout, "", 0)
	//mqtt.ERROR = log.New(os.Stdout, "", 0)
	fmt.Println("MqttBegin")

	opts := mqtt.NewClientOptions().
		AddBroker(ConfigObj.Mqtt.Server).
		SetClientID(ConfigObj.Mqtt.ClientID)
	Mqttclient = mqtt.NewClient(opts)

	if MqttToken = Mqttclient.Connect(); MqttToken.Wait() && MqttToken.Error() != nil {
		log.Println("链接失败 init")
		//panic(MqttToken.Error())
	}
	MqttToken.Wait()
	if MqttToken := Mqttclient.Subscribe(ConfigObj.Mqtt.Sub, 0, SubscribeCallback); MqttToken.Wait() && MqttToken.Error() != nil {
		fmt.Println("error?", MqttToken.Error())
	}
	go loopConntect()
}

func loopConntect() {
	for {
		time.Sleep(time.Second * 10)
		if !Mqttclient.IsConnected() {
			if MqttToken = Mqttclient.Connect(); MqttToken.Wait() && MqttToken.Error() != nil {
				//panic(MqttToken.Error())
				log.Println("重连mqtt")
			} else {
				log.Println("重连成功")
				if MqttToken := Mqttclient.Subscribe(ConfigObj.Mqtt.Sub, 0, SubscribeCallback); MqttToken.Wait() && MqttToken.Error() != nil {
					fmt.Println("error?", MqttToken.Error())
				}
			}
		}
	}

}

//发送消息
func Publish(topic string, qos byte, payload string) {
	fmt.Println("MqttPublish")
	Mqttclient.Publish(topic, qos, false, payload)
}

// sub消息回调
func SubscribeCallback(client mqtt.Client, msg mqtt.Message) {
	fmt.Println("收到消息")
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", string(msg.Payload()))
	fmt.Println("MsgId--->", msg.MessageID())

	//Publish("/device/ambit/test/pub", 2, "true")

	//encodeString := base64.StdEncoding.EncodeToString(data)
	//cache.Cache.Set(msg.Topic(), encodeString, time.Second*10)
}
