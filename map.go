package main

import "fmt"

func mp() {
	mp := make(map[string]int) //mp := map[string]int{"a": 1, "b": 2}
	mp["a"] = 1
	mp["b"] = 2

	fmt.Println(mp["a"], mp["b"])
}
