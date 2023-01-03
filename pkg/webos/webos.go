package webos

import (
	"crypto/tls"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/kaperys/go-webos"
	_ "github.com/kaperys/go-webos"
	"log"
	"net"
	"time"
)

type WebOS struct {
	TV        *webos.TV
	ClientKey string
}

func New() {
	dialer := websocket.Dialer{
		HandshakeTimeout: 10 * time.Second,
		// the TV uses a self-signed certificate
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		NetDial:         (&net.Dialer{Timeout: time.Second * 5}).Dial,
	}

	tv, err := webos.NewTV(&dialer, "<tv-ipv4-address>")
	if err != nil {
		log.Fatalf("could not dial TV: %v", err)
	}
	defer tv.Close()

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
	tv.Notification("📺👌")
}

func (w *WebOS) Volume(state bool) {
	switch state {
	case true:
		w.TV.VolumeUp()
	case false:
		w.TV.VolumeDown()
	}
}
