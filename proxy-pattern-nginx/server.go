package main

// Subject
type server interface {
	handleRequest(string, string) (int, string)
}
