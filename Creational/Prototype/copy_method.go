package main

import "fmt"

type Position struct {
	AnnualIncome     int
	Company, JobName string
}

func (p *Position) DeepCopy() *Position {
	return &Position{
		p.AnnualIncome,
		p.Company,
		p.JobName,
	}
}

type Species int

const (
	Dog = iota
	Cat
)

type Animal struct {
	Name    string
	Species Species
	Age     int
}

func (a *Animal) DeepCopy() *Animal {
	return &Animal{
		a.Name,
		a.Species,
		a.Age,
	}
}

type Human struct {
	Name     string
	Position *Position
	Animals  []*Animal
	Friends  []string
}

func (h *Human) DeepCopy() *Human {
	cpy := *h
	cpy.Position = h.Position.DeepCopy()

	for _, animal := range h.Animals {
		cpy.Animals = append(cpy.Animals, animal.DeepCopy())
	}
	copy(cpy.Friends, h.Friends)

	return &cpy
}

func main() {
	bobby := &Animal{"Bobby", Cat, 6}
	willy := &Animal{"Willy", Dog, 3}
	friends := []string{"Annette", "Wilma", "Greg"}

	john := &Human{"Max",
		&Position{
			AnnualIncome: 1000,
			Company:      "IBM",
			JobName:      "Janitor",
		},
		[]*Animal{bobby, willy},
		friends,
	}

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Friends = append(jane.Friends, "Elsa")
	jane.Animals = john.Animals[1:]

	fmt.Println(john)
	fmt.Println(jane)
}
