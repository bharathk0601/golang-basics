package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id       int    `json:"id"` // json keys in struct should be in upper case
	Name     string `json:"user-name,omitempty"`
	password string `jso:"-"`
}

func Json() {
	// marshal
	obj := Student{Id: 1, Name: "", password: "hiiii"}
	js, _ := json.Marshal(obj)
	fmt.Println(string(js))

	// unmarshal
	jsObj := `{"id": 1, "name": "bharath"}`
	jsBytes := []byte(jsObj)
	var obj2 map[string]any
	json.Unmarshal(jsBytes, &obj2)
	fmt.Println(obj2)
}
