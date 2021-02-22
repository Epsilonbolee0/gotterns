package main

import "fmt"

type Person struct {
	Name     string
	Age      int
	EyeCount int
}

func NewPerson(name string, age int) *Person {
	if age < 16 {
		panic("Age is not legal!")
	}
	return &Person{name, age, 2}
}

func main() {
	p := NewPerson("Max", 25)
	fmt.Print(p)
}
