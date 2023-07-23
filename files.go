package main

import (
	"fmt"
	"io/ioutil"
)

func readFile() {
	data, _ := ioutil.ReadFile("f.txt")
	fmt.Println(string(data))
}

func writeFile() {
	data := []byte(`hello world`)
	ioutil.WriteFile("f1.txt", data, 777)
}
func files() {
	readFile()
	writeFile()
}
