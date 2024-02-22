package connections

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	connectionsMu sync.Mutex
	connections   = make(map[*websocket.Conn]struct{})
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

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

// Send message to websocket connections
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

// Upgrade HTTP connection to websocket
func Upgrade(w http.ResponseWriter, r *http.Request) {

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error during connection upgradation:", err)
		return
	}

	Subscribe(conn)
	fmt.Println("Web Socket connected")

	defer func() {
		Unsubscribe(conn)
		conn.Close()
	}()

	for {
		// Add any additional logic needed for handling WebSocket connections here
	}
}
