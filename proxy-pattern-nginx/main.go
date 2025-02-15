package main

import (
	"fmt"
	"time"
)

func main() {
	// without proxy
	// application := &Application{}
	// application.handleRequest("/app/status", "GET")
	// application.handleRequest("/users", "POST")
	// application.handleRequest("/users", "DELETE")

	// with proxy
	proxy := NewNginxServer()

	proxy.handleRequest("/users", "POST")
	fmt.Println()
	time.Sleep(time.Second * 1)
	proxy.handleRequest("/users", "POST")
	fmt.Println()
	time.Sleep(time.Second * 1)
	proxy.handleRequest("/users", "DELETE")
	fmt.Println()
	time.Sleep(time.Second * 1)
	proxy.handleRequest("/users", "DELETE")
	fmt.Println()
	time.Sleep(time.Second * 1)

	// starts to block requests from here..
	proxy.handleRequest("/users", "POST")
	fmt.Println()
	time.Sleep(time.Second * 1)
	proxy.handleRequest("/users", "DELETE")
	fmt.Println()
	time.Sleep(time.Second * 1)
}
