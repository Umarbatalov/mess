package unit

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)


// Обработчик для веб-сокет запроса.
func WS(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)

	if err != nil {
		log.Println(err)
		return
	}

	vars := mux.Vars(r)

	if val, ok := vars["id"]; ok {
		id, _ := strconv.Atoi(val)

		client := &Client{
			id:   id,
			hub:  hub,
			conn: conn,
			send: make(chan []byte, 256),
		}

		client.hub.register <- client

		log.Println("CLient connected, id:", client.id)

		go client.WritePump()
		go client.ReadPump()
	}
}
