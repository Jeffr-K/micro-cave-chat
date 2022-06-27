package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{}

func main() {
	setRouter()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setRouter() {
	http.HandleFunc("/", home)
	http.HandleFunc("/ws", websocketRouter)
}

func home(w http.ResponseWriter, r *http.Request) {}

func websocketRouter(w http.ResponseWriter, r *http.Request) {}
