package main

import "fmt"

func sayHello(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func getHello(name string) string {
	return "Hello " + name
}

func main() {
	sayHello("Eko", "Kurniawan")
	result := getHello("Eko")
	fmt.Println(result)
}
