package subscriber

import (
	"fmt"
	mqtt "pimview.thelabshack.com/pkg/mqtt"
	"time"
)

func Run() {
	client := mqtt.GetClient("pimviewsub")

	topic := "topic/test"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic %s", topic)
	time.Sleep(time.Second * 10)

	client.Disconnect(250)
}
