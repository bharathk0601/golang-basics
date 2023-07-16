package main

import (
	"fmt"
	"strconv"
)

func typeConv() {
	b, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(b)
	}
}
