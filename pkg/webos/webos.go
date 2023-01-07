package webos

import (
	"crypto/tls"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gorilla/websocket"
	"github.com/kaperys/go-webos"
	_ "github.com/kaperys/go-webos"
	"log"
	"net"
	"strings"
	"time"
)

type WebOS struct {
	TV        *webos.TV
	ClientKey string
}

const (
	volume = "volume"
	up     = "up"
	down   = "down"
	mute   = "mute"
)

func New() *WebOS {
	dialer := websocket.Dialer{
		HandshakeTimeout: 10 * time.Second,
		// the TV uses a self-signed certificate
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		NetDial:         (&net.Dialer{Timeout: time.Second * 500}).Dial,
	}

	tv, err := webos.NewTV(&dialer, "192.168.1.36")
	if err != nil {
		log.Fatalf("could not dial TV: %v", err)
	}
	//defer tv.Close()

	// the MessageHandler must be started to read responses from the TV
	go tv.MessageHandler()

	// AuthorisePrompt shows the authorisation prompt on the TV screen
	key, err := tv.AuthorisePrompt()
	if err != nil {
		log.Fatalf("could not authorise using prompt: %v", err)
	}

	// the key returned can be used for future request to the TV using the
	// AuthoriseClientKey(<key>) method, instead of AuthorisePrompt()
	fmt.Println("Client Key:", key)

	// see commands.go for available methods
	err = tv.Notification("📺👌")
	if err != nil {
		fmt.Println(err)
	}
	return &WebOS{
		TV:        tv,
		ClientKey: key,
	}
}

// ProcessMessages process mqtt message queue and dispatch to handlers
// client mqtt broker client
// message mqtt message including topic and payload
func (w *WebOS) ProcessMessages(client mqtt.Client, message mqtt.Message) {

	fmt.Println("processing messages ...")
	fmt.Println(message.Topic())

	switch {
	case strings.Contains(message.Topic(), volume):
		w.Volume(message.Payload())
	}
}

// Volume volume handler
// direction mqtt message payload
func (w *WebOS) Volume(direction []byte) {
	d := string(direction)

	switch d {
	case up:
		fmt.Println("Volume UP")
		err := w.TV.VolumeUp()
		if err != nil {
			fmt.Println(err)
		}
	case down:
		fmt.Println("Volume DOWN")
		err := w.TV.VolumeDown()

		if err != nil {
			fmt.Println(err)
		}
	case mute:
		fmt.Println("Volume Mute")
		err := w.TV.Mute()

		if err != nil {
			fmt.Println(err)
		}
	}

}
