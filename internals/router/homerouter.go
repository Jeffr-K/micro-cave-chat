package router

import (
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("여기는 들어와짐")
	http.ServeFile(w, r, "../web")
}
