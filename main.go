package main

import (
	"log"
	"net/http"

	"github.com/Umarbatalov/mess/unit"
	"github.com/gorilla/mux"
)

func main() {
	hub := unit.NewHub()
	go hub.Run()

	r := mux.NewRouter()

	r.HandleFunc("/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		unit.WS(hub, w, r)
	})

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
