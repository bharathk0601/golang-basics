package main

import (
	"fmt"
)

// string, int, float, array, boolean -- no pointers
func changeVal(val *string) {
	*val = "hi"
}

// map, slice and func -- pointer
func changeVal2(val map[string]int) {
	val["b"] = 3
}

func ptr() {
	mp := map[string]int{"a": 1, "b": 2}
	val := "hello"
	changeVal(&val)
	fmt.Println(val)

	changeVal2(mp)
	fmt.Println(mp)
}
