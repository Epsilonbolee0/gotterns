package main

import "fmt"

type Employee struct {
	name, position string
}

type employeeMod func(employee *Employee)
type EmployeeBuilder struct {
	actions []employeeMod
}

func (b *EmployeeBuilder) Called(name string) *EmployeeBuilder {
	b.actions = append(b.actions, func(e *Employee) {
		e.name = name
	})
	return b
}

func (b *EmployeeBuilder) WorksAsA(position string) *EmployeeBuilder {
	b.actions = append(b.actions, func(e *Employee) {
		e.position = position
	})
	return b
}

func (b *EmployeeBuilder) Build() *Employee {
	person := Employee{}
	for _, action := range b.actions {
		action(&person)
	}

	return &person
}

func main() {
	b := EmployeeBuilder{}
	p := b.
		Called("Dmitri").
		WorksAsA("Developer").
		Build()

	fmt.Print(*p)
}
