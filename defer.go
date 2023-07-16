package main

import "fmt"

func print(val int) {
	fmt.Println(val)
}

func def() {
	for i := 0; i < 3; i++ {
		defer print(i)
	}
}
