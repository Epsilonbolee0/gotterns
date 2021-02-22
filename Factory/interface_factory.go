package main

import "fmt"

type Human interface {
	SayHello()
}

type human struct {
	name string
	age  int
}

type tiredHuman struct {
	name string
	age  int
}

func (h *human) SayHello() {
	fmt.Printf(" Hi, my name is %s, I am %d years old\n",
		h.name, h.age)
}

func (h *tiredHuman) SayHello() {
	fmt.Printf(" Sorry, I am to tired to talk to you")
}

func NewHuman(name string, age int) Human {
	if age > 100 {
		return &tiredHuman{name, age}
	}
	return &human{name, age}
}

func main() {
	h := NewHuman("Sashka", 134)
	h.SayHello()
}
