package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrater = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(string(p))
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}
func servews(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)
	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrater.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	reader(ws)
}
func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "simple server")
	})
	http.HandleFunc("/ws", servews)
	log.Println("Starting Server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
func main() {
	setupRoutes()
	fmt.Println("Chat App v0.01")
}
