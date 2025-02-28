package main

import "fmt"

type Blacklst func(string) bool

func registerUser(name string, blacklist Blacklst) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "Anjing"
	}

	registerUser("Eko", blacklist)
	registerUser("Anjing", func(name string) bool {
		return name == "Anjing"
	})
}
