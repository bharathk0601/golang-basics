package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	id   int
	name string
}

func jsonDt() {
	obj := student{id: 1, name: "bharath"}
	js, _ := json.Marshal(obj)
	fmt.Println(js)

}
