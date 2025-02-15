package main

func main() {
	application := &Application{}
	application.handleRequest("/app/status", "GET")
	application.handleRequest("/users", "POST")
	application.handleRequest("/users", "DELETE")
}
