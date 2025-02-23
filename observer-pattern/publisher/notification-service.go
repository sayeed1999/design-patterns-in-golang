package publisher

import "github.com/sayeed1999/design-patterns-in-golang/observer-pattern/observer"

// implements Subject
type NotificationService struct {
	observers []observer.Observer
}

func (s *NotificationService) NotifyObservers() {
	for _, ob := range s.observers {
		ob.Update()
	}
}

func (s *NotificationService) AddObserver(observer observer.Observer) {
	s.observers = append(s.observers, observer)
}

func (s *NotificationService) RemoveObserver(observer observer.Observer) {
	for i, value := range s.observers {
		if value == observer {
			s.observers = append(s.observers[0:i], s.observers[i+1:]...)
			break
		}
	}
}
