package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Speed struct {
	Speed int `json:"speed"`
}

func main() {
	mqttOpts := mqtt.NewClientOptions()
	mqttOpts.AddBroker("localhost:1883")
	mqttOpts.SetClientID("test-client")
	client := mqtt.NewClient(mqttOpts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Println("error occured creating the mqtt client:", token.Error())
	}
	rand.Seed(time.Now().UnixNano())
	log.Println("publishing to topic topic/test")
	for {
		speed := rand.Intn(100) + 1
		payload, _ := json.Marshal(Speed{Speed: speed})
		client.Publish("topic/test", 0, false, payload)
		log.Println("puslished message with speed:", speed)
		time.Sleep(1 * time.Second)
	}
}
