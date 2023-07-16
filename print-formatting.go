package main

import "fmt"

func printFormatting() {
	age := 25
	name := "bharath"
	sal := 25.55

	fmt.Println("hello world")
	fmt.Print("hello \nworld \n")
	fmt.Println("my age is", age)
	fmt.Printf("my age is %v and name is %v\n", age, name)
	fmt.Printf("my age is %v and name is %q\n", age, name)
	fmt.Printf("age is of type %T\n", age)
	fmt.Printf("sal is %0.2f\n", sal)
}
