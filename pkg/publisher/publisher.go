package publisher

import (
	"fmt"
	mqtt "pimview.thelabshack.com/pkg/mqtt"
	"time"
)

func Publish() {
	client := mqtt.GetClient("pimviewpub")

	num := 10
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("Message %d", i)
		token := client.Publish("topic/test", 0, false, text)
		token.Wait()
		time.Sleep(time.Second)
	}
}
