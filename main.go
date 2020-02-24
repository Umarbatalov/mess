package main

import (
	"log"
	"net/http"

	"github.com/Umarbatalov/mess/node"
	"github.com/gorilla/mux"
)

func main() {
	hub := node.NewHub()
	go hub.Run()

	r := mux.NewRouter()

	r.HandleFunc("/ws/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		node.WebSocketHandler(hub, w, r)
	})

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
