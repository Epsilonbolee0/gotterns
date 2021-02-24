package main

import "fmt"

type Worker struct {
	Name, Position string
	AnnualIncome   int
}

const (
	Developer = iota
	Manager
)

func NewWorker(role int) *Worker {
	switch role {
	case Developer:
		return &Worker{"", "Developer", 60000}
	case Manager:
		return &Worker{"", "Manager", 80000}
	default:
		panic("Unsupported role!")
	}
}

func main() {
	m := NewWorker(Manager)
	m.Name = "Sam"
	fmt.Println(m)
}
