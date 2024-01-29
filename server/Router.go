package main

import "sync"

type Router struct {
	mu      sync.Mutex
	clients []*Client
}

func newRouter() *Router {
	return &Router{
		clients: make([]*Client, 0),
	}
}
