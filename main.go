package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"real-time-chat/controllers"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Dotenv not loaded: ", err)
	}
}

func main() {
	port := os.Getenv("PORT")

	room := controllers.NewRoom()

	go room.Run()

	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/ws", func(rw http.ResponseWriter, r *http.Request) {
		controllers.Server(room, rw, r)
	})

	log.Printf("Starting server on the port %v...", port)

	if len(port) == 0 {
		log.Fatal("DotEnv does load")
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
