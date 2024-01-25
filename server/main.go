package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"golang.org/x/net/websocket"
)

type Connection struct {
	id int
	ws *websocket.Conn
}

// array of connections
var connections []Connection

var freeId int = 0

func Echo(ws *websocket.Conn) {
	var err error

	var connectionId = freeId
	freeId++

	//add connection to array
	connections = append(connections, Connection{connectionId, ws})

	for {
		var reply string

		err = websocket.Message.Receive(ws, &reply)

		if err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)

		msg := "Received:  " + reply + " from connection " + strconv.Itoa(connectionId)
		fmt.Println("Sending to client: " + msg)

		//for each connection in array, send message
		for _, conn := range connections {
			if conn.id != connectionId {
				err = websocket.Message.Send(conn.ws, msg)

				if err != nil {
					fmt.Println("Can't send")
					break
				}
			}
		}

	}
}

func main() {
	print("Server started")
	http.Handle("/", websocket.Handler(Echo))

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
