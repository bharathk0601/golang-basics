package main

import "fmt"

// array is fixed and slices can be changed
func array() {
	// var ages [3]int = [3]int{10, 20, 30}
	var ages = [3]int{10, 20, 30}
	names := [4]string{"manoj", "bharath", "shettigar"}
	names[1] = "manoj kumar"
	fmt.Println(ages, len(ages), names, len((names)))

	var scores = []int{100, 50, 60}
	scores[2] = 10
	scores = append(scores, 30)
	fmt.Println(scores, len(scores))

	range1 := names[0:2]
	range2 := names[2:]
	fmt.Println(range1, range2)
}
