package main

import "sync"

// Router keeps track of all clients.
// also offers a mutex so access is safe
type Router struct {
	mu      sync.Mutex
	clients []*Client
}

// returns a new Router
func newRouter() *Router {
	return &Router{
		clients: make([]*Client, 0),
	}
}
