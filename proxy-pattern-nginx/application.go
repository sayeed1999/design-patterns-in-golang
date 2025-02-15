package main

import "fmt"

// Real Subject
type Application struct{}

func (app *Application) handleRequest(url, method string) (int, string) {
	var status int
	var message string

	if url == "/app/status" && method == "GET" {
		fmt.Printf("[Application][Success]: handling request '%s %s'\n", method, url)
		status = 200
		message = "Ok"
	} else if url == "/users" && method == "POST" {
		fmt.Printf("[Application][Success]: handling request '%s %s'\n", method, url)
		status = 201
		message = "User created"
	} else {
		fmt.Println(fmt.Errorf("[Application][Error]: Could not handle request '%s %s'", method, url))
		status = 404
		message = "Endpoint not found!"
	}

	return status, message
}
