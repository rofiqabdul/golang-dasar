package main

import "fmt"

func main() {
	name := "Abdul Rofiq"
	fmt.Println(name)

	name = "Rofiq Abdul"
	fmt.Println(name)

	var (
		firstname  = "Abdul"
		middleName = "Rofiq"
		lastName   = "Hidayat"
	)

	fmt.Println(firstname + " " + middleName + " " + lastName)
}
