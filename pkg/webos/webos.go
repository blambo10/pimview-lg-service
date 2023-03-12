package webos

import (
	"crypto/tls"
	"fmt"
	webos "github.com/blambo10/go-webos"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gorilla/websocket"
	"log"
	"net"
	"pimview.thelabshack.com/pkg/config"
	"strings"
	"time"
)

//"github.com/kaperys/go-webos"
//_ "github.com/kaperys/go-webos"

type WebOS struct {
	TV        *webos.TV
	ClientKey string
	Host      string
}

const (
	volume = "volume"
	up     = "up"
	down   = "down"
	mute   = "mute"
)

var (
	cfg = config.GetWebOS()
)

func New() (*WebOS, error) {

	//todo: finish connection logic, abstract it so that its done on instantiation and also if not connected at call to tv,
	//      recall connect, fail graceful not fatal if connection failed.

	//tv, key, err := tv.Connect()
	//if err != nil {
	//	return nil, err
	//}
	dialer := websocket.Dialer{
		HandshakeTimeout: 10 * time.Second,
		// the TV uses a self-signed certificate
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		NetDial:         (&net.Dialer{Timeout: time.Second * 500}).Dial,
	}

	tv, err := webos.NewTV(&dialer, cfg.Host)
	if err != nil {
		fmt.Println("could not dial TV: %v", err)
		return nil, fmt.Errorf("cannot connect to tv")
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
		Host:      cfg.Host,
	}, nil
}

func (w *WebOS) Connect() (*webos.TV, string, error) {
	var tv *webos.TV
	var key string

	dialer := websocket.Dialer{
		HandshakeTimeout: 10 * time.Second,
		// the TV uses a self-signed certificate
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		NetDial:         (&net.Dialer{Timeout: time.Second * 500}).Dial,
	}

	tv, err := webos.NewTV(&dialer, cfg.Host)
	if err != nil {
		fmt.Println("could not dial TV: %v", err)
		return tv, key, fmt.Errorf("cannot connect to tv")
	}
	//defer tv.Close()

	// the MessageHandler must be started to read responses from the TV
	go tv.MessageHandler()

	// AuthorisePrompt shows the authorisation prompt on the TV screen
	key, err = tv.AuthorisePrompt()
	if err != nil {
		return tv, key, fmt.Errorf("cannot connect to tv")
	}

	// the key returned can be used for future request to the TV using the
	// AuthoriseClientKey(<key>) method, instead of AuthorisePrompt()
	fmt.Println("Client Key:", key)

	// see commands.go for available methods
	err = tv.Notification("📺👌")
	if err != nil {
		fmt.Println(err)
		//return fmt.Errorf("cannot connect to tv")
	}

	return tv, key, nil
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
func (w *WebOS) Volume(direction []byte) error {
	d := string(direction)
	fmt.Println("client key:", w.ClientKey)
	switch d {
	case up:
		fmt.Println("Volume UP")
		err := w.TV.VolumeUp()
		if err != nil {
			fmt.Println(err)
			return err
		}
	case down:
		fmt.Println("Volume DOWN")
		err := w.TV.AuthoriseClientKey(w.ClientKey)
		if err != nil {
			fmt.Println(err)
			return err
		}

		err = w.TV.VolumeDown()

		if err != nil {
			fmt.Println(err)
			return err
		}
	case mute:
		fmt.Println("Volume Mute")
		err := w.TV.Mute()

		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil

}
