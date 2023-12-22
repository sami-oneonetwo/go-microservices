package connections

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/sami-oneonetwo/go-microservices/subscriber/messages"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Upgrade(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error during connection upgradation:", err)
		return
	}

	messages.Subscribe(conn)
	fmt.Println("Web Socket connected")

	defer func() {
		messages.Unsubscribe(conn)
		conn.Close()
	}()

	for {
		// Add any additional logic needed for handling WebSocket connections here
	}
}
