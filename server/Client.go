package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"golang.org/x/net/websocket"
)

type Client struct {
	id     int
	conn   *websocket.Conn
	router *Router
	send   chan string
}

// creates a client and registers it at the router
// locks the routers clients so it is safe to append
func newClient(id int, conn *websocket.Conn, router *Router) *Client {
	client := Client{
		id,
		conn,
		router,
		make(chan string),
	}

	router.mu.Lock()
	router.clients = append(router.clients, &client)
	router.mu.Unlock()

	return &client
}

// runs booth functions in go routines, so they run in parallel
// also puts them in a wait group to keep them alive.
func (c *Client) run() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		c.sendMessage()
	}()

	go func() {
		defer wg.Done()
		c.receiveMessage()
	}()

	wg.Wait()
}

// if it receives a message from the client it sends it to the recipient
func (c *Client) receiveMessage() {
	var err error
	var message string
	for {

		err = websocket.Message.Receive(c.conn, &message)

		if err != nil {
			fmt.Println("Can't receive", err)
			break
		}

		//TODO make the message processing more sophisticated
		//a message looks like this: "recipientId;messageContent"
		parsedMessage := strings.SplitN(message, ";", 2)

		recipientId, conversionErr := strconv.Atoi(parsedMessage[0])

		if conversionErr != nil {
			fmt.Println("couldn't convert client id")
			continue
		}

		messageContent := parsedMessage[1]

		//locks it so its save to access the clients
		c.router.mu.Lock()
		for _, client := range c.router.clients {
			if client.id == recipientId {
				client.send <- messageContent
				c.send <- messageContent
			}
		}
		c.router.mu.Unlock()
	}
}

// if it gets a message over the chanel it sends it to the connected client
func (c *Client) sendMessage() {
	var err error
	var message string

	for {
		message = <-c.send
		err = websocket.Message.Send(c.conn, message)

		if err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}
