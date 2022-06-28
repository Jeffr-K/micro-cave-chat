package router

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func WebsocketRouter(w http.ResponseWriter, r *http.Request) {
	log.Println("Request: ", r.Header)
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	websocket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebsocketRouter", err)
	}
	defer websocket.Close()
	log.Println("Client Successfully Connected..")
	reader(websocket)
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(p))
		message := []byte("Let's start to talk with me")
		if err := conn.WriteMessage(messageType, message); err != nil {
			log.Println(err)
			return
		}
	}
}
