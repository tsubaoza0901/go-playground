package main

import "fmt"

type Message string

func NewMessage() Message {
	return Message("")
}

func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

type Greeter struct {
	Message Message
}

func (g Greeter) Greet() Message {
	return g.Message
}

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

type Event struct {
	Greeter Greeter
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func main() {
	// -------------------
	// // without using Wire
	// message := NewMessage()
	// greeter := NewGreeter(message)
	// event := NewEvent(greeter)

	// event.Start()
	// -------------------

	// -------------------
	// // using Wire
	e := InitializeEvent()

	e.Start()
	// -------------------
}
