package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Abdul"
	names[1] = "Rofiq"
	names[2] = "Hidayat"

	fmt.Println(names)

	var values = [...]int{
		90,
		80,
		95,
	}

	fmt.Println(values)
	fmt.Println(len(values))
	values[0] = 100
	fmt.Println(values)
}
