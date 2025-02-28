package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := newMap("")

	if data == nil {
		fmt.Println("Data map kosong")
	} else {
		fmt.Println(data)
	}
}
