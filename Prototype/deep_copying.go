package main

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
}

func main() {
	sashka := Person{"Sahka",
		&Address{"Red Torch 1", "Novosibirsk", "Russia"}}

	dashka := sashka
	dashka.Address = &Address{
		sashka.Address.StreetAddress,
		"Moscow",
		sashka.Address.Country,
	}

	fmt.Println(sashka, sashka.Address)
	fmt.Println(dashka, dashka.Address)
}
