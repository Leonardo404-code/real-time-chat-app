package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func Server(room *Room, w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	user := &User{room: room, conn: ws, send: make(chan []byte, 256)}

	user.room.register <- user

	log.Println("user connected")

	go user.writePump()
	go user.readPump()
}
