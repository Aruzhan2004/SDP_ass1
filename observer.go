package main

import (
	"fmt"
)

type Observable interface {
	Add(observer Observer)
	Remove(observer Observer)
	NotifyObservers()
}

// Observer is the interface for the observers.
type Observer interface {
	Update(value int)
}

// NewsResource implements the Observable interface.
type NewsResource struct {
	value     int
	observers []Observer
}

func (nr *NewsResource) Add(observer Observer) {
	nr.observers = append(nr.observers, observer)
}

func (nr *NewsResource) Remove(observer Observer) {
	for i, obs := range nr.observers {
		if obs == observer {
			nr.observers = append(nr.observers[:i], nr.observers[i+1:]...)
			break
		}
	}
}

func (nr *NewsResource) NotifyObservers() {
	for _, observer := range nr.observers {
		observer.Update(nr.value)
	}
}

// NewsAgency implements the Observer interface.
type NewsAgency struct{}

func (na *NewsAgency) Update(value int) {
	fmt.Printf("NewsAgency handles updated value: %d\n", value)
}

// Reporter implements the Observer interface.
type Reporter struct{}

func (r *Reporter) Update(value int) {
	fmt.Printf("Reporter. Updated value: %d\n", value)
}

// Blogger implements the Observer interface.
type Blogger struct{}

func (b *Blogger) Update(value int) {
	fmt.Printf("Blogger. New article about value: %d\n", value)
}

func main() {
	resource := &NewsResource{}
	newsAgency := &NewsAgency{}
	reporter := &Reporter{}
	blogger := &Blogger{}

	resource.Add(newsAgency)
	resource.NotifyObservers()
	fmt.Println("NewsAgency added value")
	resource.value = 1

	resource.Add(reporter)
	resource.NotifyObservers()
	fmt.Println("Reporter added value")
	resource.value = 4

	resource.Add(blogger)
	resource.NotifyObservers()
	fmt.Println("Blogger added value")
	resource.value = 7
}
