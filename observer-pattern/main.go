package main

import (
	"github.com/sayeed1999/design-patterns-in-golang/observer-pattern/observer"
	"github.com/sayeed1999/design-patterns-in-golang/observer-pattern/publisher"
)

func main() {
	email := &observer.Email{}
	sms := &observer.SMS{}

	subject := &publisher.NotificationService{}
	subject.AddObserver(email)
	subject.AddObserver(sms)

	subject.NotifyObservers()
}
