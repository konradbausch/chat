package main

import (
	"fmt"
	"golang.org/x/net/websocket"
)

type Client struct {
	id     string
	conn   *websocket.Conn
	router *Router
	sent   chan string
}

func (c *Client) newClient(id string, conn *websocket.Conn, router *Router) *Client {
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

func (c *Client) run() {
	go c.sentMessage()
	go c.receiveMessage()
}

func (c *Client) receiveMessage() {
	var err error
	var message string
	for {

		err = websocket.Message.Receive(c.conn, &message)

		if err != nil {
			fmt.Println("Can't receive")
			break
		}

		//TODO handle sent
		fmt.Println("Still to do")

	}
}

func (c *Client) sentMessage() {
	var err error
	var message string

	for {
		message = <-c.sent
		err = websocket.Message.Send(c.conn, message)

		if err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}
