package main

import (
	"fmt"
)

type Person struct {
	// Address info
	StreetAddress, Postcode, City string

	// Job info
	CompanyName, Position string
	AnnualIncome          int

	// Private Info
	Name, Surname string
	Age           int
}

type PersonBuilder struct {
	person *Person
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

func (b *PersonBuilder) Has() *PersonPrivateBuilder {
	return &PersonPrivateBuilder{*b}
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

type PersonAddressBuilder struct {
	PersonBuilder
}

func (it *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	it.person.StreetAddress = streetAddress
	return it
}

func (it *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	it.person.City = city
	return it
}

func (it *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	it.person.Postcode = postcode
	return it
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (it *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	it.person.CompanyName = companyName
	return it
}

func (it *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	it.person.Position = position
	return it
}

func (it *PersonJobBuilder) Earning(income int) *PersonJobBuilder {
	it.person.AnnualIncome = income
	return it
}

type PersonPrivateBuilder struct {
	PersonBuilder
}

func (it *PersonPrivateBuilder) Name(name string) *PersonPrivateBuilder {
	it.person.Name = name
	return it
}

func (it *PersonPrivateBuilder) Surname(surname string) *PersonPrivateBuilder {
	it.person.Surname = surname
	return it
}

func (it *PersonPrivateBuilder) Age(age int) *PersonPrivateBuilder {
	it.person.Age = age
	return it
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}

func main() {
	pb := NewPersonBuilder()
	pb.
		Lives().
		At("Mayakovskogo 17").
		In("Novosibirsk").
		WithPostcode("14881337").
		Works().
		At("RZHD").
		AsA("Locksmith").
		Earning(30000).
		Has().
		Name("Ivan").
		Surname("Govnov").
		Age(27)

	person := pb.Build()
	fmt.Println(person)
}
