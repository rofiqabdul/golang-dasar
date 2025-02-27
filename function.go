package main

import "fmt"

func sayHello(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func getHello(name string) string {
	return "Hello " + name
}

func getFullName() (string, string) {
	return "Eko", "Khaneedy"
}

func main() {
	sayHello("Eko", "Kurniawan")
	result := getHello("Eko")
	fmt.Println(result)

	fmt.Println(getFullName())
	firstName, _ := getFullName()
	fmt.Println(firstName)
}
