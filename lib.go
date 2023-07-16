package main

import (
	"fmt"
	"sort"
	"strings"
)

func lib() {
	greetings := "hello world"
	fmt.Println(strings.Contains(greetings, "hello"))
	fmt.Println(strings.ReplaceAll(greetings, "hello", "hi"))
	fmt.Println(strings.ToUpper((greetings)))
	fmt.Println(strings.Index(greetings, "ll"))
	fmt.Println(strings.Split(greetings, " "))
	fmt.Println(greetings[3:])

	ages := []int{9, 4, 6, 8, 11}
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 6)
	fmt.Println(index)

	names := []string{"manoj", "bharath"}
	sort.Slice(names, func(i, j int) bool {
		return names[i] > names[j]
	})
	fmt.Println(names)
}
