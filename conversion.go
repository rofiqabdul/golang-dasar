package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name = "Rofiq"
	var e = name[0]
	var eStirng = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eStirng)
}
