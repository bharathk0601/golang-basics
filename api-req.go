package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func makeApiCall() {
	makePostReq()
}

func makeGetReq(url string) {
	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	resBody, _ := ioutil.ReadAll(res.Body)

	var js []map[string]interface{}
	json.Unmarshal(resBody, &js)

	fmt.Println(url, "StatusCode", res.StatusCode)
}

func makePostReq() {
	const url = "https://jsonplaceholder.typicode.com/posts"
	body := map[string]interface{}{"title": "Post title", "body": "Post description", "userId": 1}
	jsString, _ := json.Marshal(body)
	res, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(jsString)))

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	resBody, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(resBody))
}
