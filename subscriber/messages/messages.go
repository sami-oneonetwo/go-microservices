package messages

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	connectionsMu sync.Mutex
	connections   = make(map[*websocket.Conn]struct{})
)

func Subscribe(conn *websocket.Conn) {
	connectionsMu.Lock()
	defer connectionsMu.Unlock()

	connections[conn] = struct{}{}
}

func Unsubscribe(conn *websocket.Conn) {
	connectionsMu.Lock()
	defer connectionsMu.Unlock()

	delete(connections, conn)
}

func SendMessage(message string) {
	connectionsMu.Lock()
	defer connectionsMu.Unlock()

	fmt.Println("Pushing to web sockets")
	for conn := range connections {
		err := conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			fmt.Println("Error during message writing:", err)
		}
	}
}
