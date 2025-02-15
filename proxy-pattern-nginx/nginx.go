package main

import "fmt"

// Proxy
type Nginx struct {
	application       *Application
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func NewNginxServer() *Nginx {
	return &Nginx{
		application:       &Application{},
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}

func (n *Nginx) handleRequest(url, method string) (int, string) {

	allowed := n.checkRateLimitting(method, url)

	if !allowed {
		fmt.Println(fmt.Errorf("[Proxy][Error]: Blocked request '%s %s'", method, url))
		return 403, "Not Allowed!"
	}

	fmt.Printf("[Proxy][Success]: Proxying request to application '%s %s'\n", method, url)
	return n.application.handleRequest(url, method)
}

func (n *Nginx) checkRateLimitting(method, url string) bool {
	count := n.rateLimiter[method+url]

	if count >= n.maxAllowedRequest {
		return false
	}

	n.rateLimiter[method+url]++
	return true
}
