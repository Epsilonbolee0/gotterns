package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Material int

const (
	Steel = iota
	Zirconium
	Carbon
)

type Engine struct {
	Volume, Pistons int
	Material        Material
}

type Country int

const (
	Russia = iota
	America
	Sweden
)

type Car struct {
	Brand       string
	Distributor Country
	Engine      *Engine
	Owners      []string
}

func (c *Car) DeepCopy() *Car {
	buffer := bytes.Buffer{}
	encoder := gob.NewEncoder(&buffer)
	_ = encoder.Encode(c)

	decoder := gob.NewDecoder(&buffer)
	result := Car{}
	_ = decoder.Decode(&result)
	return &result
}

func main() {
	volvo := &Car{
		Brand:       "Volvo",
		Distributor: Russia,
		Engine: &Engine{
			Volume:   80,
			Pistons:  4,
			Material: Zirconium,
		},
		Owners: []string{"Vanya", "Oleg", "Denis", "Igor"},
	}

	gasel := volvo.DeepCopy()
	gasel.Brand = "Gasel"
	gasel.Engine.Volume = 50
	gasel.Owners = append(gasel.Owners[1:3], "Stas")

	fmt.Println(volvo)
	fmt.Println(gasel)
}
