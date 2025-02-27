package main

import "fmt"

func main() {
	for counter := 1; counter <= 10; counter++ {
		if counter == 5 {
			break
		}
		fmt.Println("Perulangan ke", counter)
	}

	names := []string{"Eko", "Kurniawan", "Khannedy"}

	for i := 0; i < len(names); i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(names[i])
	}

	for index, name := range names {
		fmt.Println("index", index, "name", name)
	}

	for _, name := range names {
		fmt.Println("name", name)
	}
}
