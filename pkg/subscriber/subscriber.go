package subscriber

import (
	"fmt"
	pahomqtt "github.com/eclipse/paho.mqtt.golang"
	mqtt "pimview.thelabshack.com/pkg/mqtt"
	webos "pimview.thelabshack.com/pkg/webos"
	"time"
)

func Run() {
	webos := webos.New()
	client := mqtt.GetClient("pimview")

	for {
		//Sub to mqtt topic (clean up later)
		Subscribe(client, webos.ProcessMessages)
	}
}

func Subscribe(cc pahomqtt.Client, handler pahomqtt.MessageHandler) {
	topic := "webos/volume"
	token := cc.Subscribe(topic, 1, handler)
	token.Wait()
	fmt.Printf("Subscribed to topic %s", topic)
	time.Sleep(time.Second * 120)
}
