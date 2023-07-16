package main

import "fmt"

func cond() {
	age := 25

	if age < 25 {
		fmt.Println("age < 25")
	} else if age < 20 {
		fmt.Println("age < 20")
	} else {
		fmt.Println("age is less...")
	}

	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}

		fmt.Println(i)
	}
}
