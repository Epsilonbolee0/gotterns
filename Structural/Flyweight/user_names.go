package main

import (
	"fmt"
	"strings"
)

type User struct {
	FullName string
}

func NewUser(fullName string) *User {
	return &User{FullName: fullName}
}

var allNames []string

type BetterUser struct {
	names []uint16
}

func NewBetterUser(fullName string) *BetterUser {
	getOrAdd := func(s string) uint16 {
		for i := range allNames {
			if allNames[i] == s {
				return uint16(i)
			}
		}
		allNames = append(allNames, s)
		return uint16(len(allNames) - 1)
	}

	result := BetterUser{}
	parts := strings.Split(fullName, " ")
	for _, p := range parts {
		result.names = append(result.names, getOrAdd(p))
	}

	return &result
}

func (u *BetterUser) FullName() string {
	var parts []string
	for _, id := range u.names {
		parts = append(parts, allNames[id])
	}

	return strings.Join(parts, " ")
}

func main() {
	john := NewBetterUser("John Doe")
	//jane := NewBetterUser("Jane Doe")
	//alsoJane := NewBetterUser("Jane Smith")

	fmt.Println(john.FullName())

}
