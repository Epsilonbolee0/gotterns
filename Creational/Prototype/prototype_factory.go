package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type WorkAddress struct {
	Suite               int
	StreetAddress, City string
}

type Employee struct {
	Name   string
	Office WorkAddress
}

func (e *Employee) DeepCopy() *Employee {
	buffer := bytes.Buffer{}
	encoder := gob.NewEncoder(&buffer)
	_ = encoder.Encode(e)

	d := gob.NewDecoder(&buffer)
	result := Employee{}
	_ = d.Decode(&result)
	return &result
}

var mainOffice = Employee{
	"", WorkAddress{0, "123 Lubyanka", "Moscow"},
}

var auxOffice = Employee{
	"", WorkAddress{0, "42 Basmannaya", "Moscow"},
}

func NewEmployee(proto *Employee, name string, suite int) *Employee {
	result := proto.DeepCopy()
	result.Name = name
	result.Office.Suite = suite
	return result
}

func NewMainOfficeEmployee(name string, suite int) *Employee {
	return NewEmployee(&mainOffice, name, suite)
}

func NewAuxOfficeEmployee(name string, suite int) *Employee {
	return NewEmployee(&auxOffice, name, suite)
}

func main() {
	john := NewAuxOfficeEmployee("John", 100)
	jane := NewMainOfficeEmployee("Jane", 5)

	fmt.Println(john)
	fmt.Println(jane)
}
