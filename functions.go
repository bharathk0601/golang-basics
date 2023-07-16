package main

import (
	"fmt"
	"math"
)

func cycleNames(names []string, f func(string)) {
	for _, v := range names {
		f(v)
	}
}

func getAreaOfCircle(r float64) float64 {
	return math.Phi * r * r
}

func mulReturn() (string, int) {
	return "hello", 1
}

func funcTest() {
	cycleNames([]string{"bharath", "manoj"}, printStr)
	// a1 := getAreaOfCircle(15)
	val1, val2 := mulReturn()
	fmt.Println(val1, val2)
}
