package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	fmt.Println("Hello World")

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.OnConnect("connection", func(so socketio.Socket) {
		log.Println("New Connection")

		so.Join("chat")

		so.On("chat message", func(msg string) {
			log.Println("Message Received From Client: " + msg)
			so.BroadcastTo("chat", "chat message", msg)
		})
	})

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.Handle("/socket.io", server)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
