package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil function")
}

func runApplication() {
	defer logging()
	fmt.Println("Run application")
}

func endApp() {
	fmt.Println("End app")
	message := recover()
	fmt.Println("Terjadi error", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Ups Error")
	}
}

func main() {
	runApplication()
	runApp(true)
}
