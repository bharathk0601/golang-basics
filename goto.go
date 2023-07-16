package main

import "fmt"

func goTo() {
	fmt.Println("hi")
	fmt.Println("1")
	goto lablel
	fmt.Println("2")
lablel:
	fmt.Println("3")
	fmt.Println("4")

}
