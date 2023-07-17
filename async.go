package main

import (
	"fmt"
	"time"
)

func fun1() {
	time.Sleep(10 * 5)
}

func fun2() {
	time.Sleep(5 * 5)
}

func async() {
	go fun1()
	fmt.Println("hello")
	go fun2()
}
