package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var c = a*b - 100/2 + 10%2

	fmt.Println(c)

	var i = 10
	i += 10

	fmt.Println(i)
}
