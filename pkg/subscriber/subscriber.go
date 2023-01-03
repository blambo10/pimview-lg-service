package subscriber

import (
	"fmt"
	mqtt "pimview.thelabshack.com/pkg/mqtt"
	"time"
)

// TODO: Intermittent mtqq from websockets being received, see why (its running on localhost)
func Run() {
	client := mqtt.GetClient("pimview-webos-sub")

	topic := "topic/test"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic %s", topic)
	time.Sleep(time.Second * 30)

	client.Disconnect(250)
}

//func Run() {
//
//	//ticker := time.NewTicker(30 * time.Second)
//	//done := make(chan struct{})
//	//
//	//for {
//	client := mqtt.GetClient("pimview-webos-sub")
//	//Subscribe(client)
//	topic := "topic/test"
//	token := client.Subscribe(topic, 1, nil)
//	token.Wait()
//	fmt.Printf("Subscribed to topic %s", topic)
//	time.Sleep(time.Second * 10)
//	client.Disconnect(250)
//	//
//	//	select {
//	//	case <-done:
//	//		return
//	//	case <-ticker.C:
//	//	}
//	//}
//
//	//time.Sleep(time.Second * 30)
//
//}
//
//func Subscribe(c *pahomqtt.Client) {
//
//	//client := mqtt.GetClient("pimviewsub")
//
//	//topic := "topic/test"
//	//token := client.Subscribe(topic, 1, nil)
//	//token.Wait()
//	//fmt.Printf("Subscribed to topic %s", topic)
//	//time.Sleep(time.Second * 10)
//	//
//	//client.Disconnect(250)
//}
