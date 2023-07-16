package main

import "fmt"

func loop() {
	x := 0
	for x < 5 {
		fmt.Println(x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	names := []string{"bharath", "manoj", "shreys"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for k, v := range names {
		fmt.Println(k, v)
	}

	towd := [][]int{{1, 2}, {2, 3}}
	fmt.Println((towd))

	for _, v1 := range towd {
		for k, v := range v1 {
			fmt.Println(k, v)
		}
	}

	dynamic := []any{1, 2, 3, "hello"}
	fmt.Println(dynamic)
}
