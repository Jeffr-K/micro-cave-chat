package router

import "net/http"

func SetRouter() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/ws", WebsocketRouter)
}
