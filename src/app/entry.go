package main

import (
	"fmt"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/pixeljonas/frontipi/src/config"
	"github.com/pixeljonas/frontipi/src/mqtt"
)

var handler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

var commandHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	mqtt.ApplyCommand(string(msg.Payload()))
	mqtt.PublishState(client)
}

func statePublisher(client MQTT.Client, interval int) {
	timer := time.NewTicker(time.Duration(interval) * time.Minute)
	for t := range timer.C {
		fmt.Println("[State Publisher] posting state", t)
		mqtt.PublishState(client)
	}
}

func main() {

	client := mqtt.Connect()
	cfg := config.Get()
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	msg := fmt.Sprintf("%s/%s", cfg.RootTopic, cfg.CommandTopic)
	fmt.Println(msg)
	if token := client.Subscribe(fmt.Sprintf("%s/%s", cfg.RootTopic, cfg.CommandTopic), 0, commandHandler); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	mqtt.AutoDiscover(client)

	go statePublisher(client, cfg.StateInterval)

	timer := time.NewTicker(1 * time.Minute)
	for t := range timer.C {
		fmt.Println("running", t)
	}
}
