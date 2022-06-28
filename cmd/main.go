package main

import (
	"fmt"
	"log"
	"micro-cave-chat/internals/router"
	"net/http"
)

func main() {
	router.SetRouter()
	fmt.Println("Cave chat module is running on 6700..")
	log.Fatal(http.ListenAndServe(":6700", nil))
}
