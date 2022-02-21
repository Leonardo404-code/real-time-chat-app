package main

import (
	"fmt"
	"log"
	"net/http"
	"real-time-chat/controllers"
)

func main() {
	room := controllers.NewRoom()

	go room.Run()

	http.HandleFunc("/ws", func(rw http.ResponseWriter, r *http.Request) {
		controllers.Server(room, rw, r)
	})

	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
