package controllers

import (
	"bytes"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait = 10 * time.Second

	pongWait = 60 * time.Second

	pingPeriod = (pongWait * 9) / 10
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// User is a middleman between the websocket connection and the room.
type User struct {
	room *Room

	conn *websocket.Conn

	send chan []byte
}

func (u *User) readPump() {
	defer func() {
		u.room.unregister <- u
		u.conn.Close()
	}()

	for {
		_, message, err := u.conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Fatalf("error: %v", err)
			}
			break
		}

		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))

		u.room.broadcast <- message
	}
}

func (u *User) writePump() {
	defer func() {
		u.conn.Close()
	}()

	for {
		select {
		case message, ok := <-u.send:
			if !ok {
				u.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := u.conn.NextWriter(websocket.TextMessage)

			if err != nil {
				return
			}

			w.Write(message)

			n := len(u.send)

			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-u.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}
