package publisher

import "github.com/sayeed1999/design-patterns-in-golang/observer-pattern/observer"

type Subject interface {
	// need to have observers[]
	NotifyObservers()
	AddObserver(observer.Observer)
	RemoveObserver(observer.Observer)
}
