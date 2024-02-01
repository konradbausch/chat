package main

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

// the next unique id for a user
var freeId = 0

// keeps track of all clients
var router = newRouter()

func main() {
	print("Server started")
	http.Handle("/", websocket.Handler(newConnection))

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

// creates a client and runs it
func newConnection(conn *websocket.Conn) {
	client := newClient(freeId, conn, router)
	freeId++
	client.run()

}
