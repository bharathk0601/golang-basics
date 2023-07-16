package main

import "fmt"

func num() {
	var name1 string = "Bharath"
	var name2 = "Manoj"
	var name3 string
	name4 := "bharath2"

	fmt.Println("hello world!", name1, name2, name3, name4)

	var age1 int = 20
	var age2 int = 30 // default int32
	age3 := 40
	var num1 int8 = 127
	var num2 uint8 = 255
	var num3 float32 = 1.25
	var num4 float64 = 1.24
	num5 := 1.5

	fmt.Println(age1, age2, age3, num1, num2, num3, num3, num4, num5)
}
