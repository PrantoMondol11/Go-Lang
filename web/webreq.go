package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.google.com"

func main() {

	fmt.Println("Web request")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	databyte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databyte)
	fmt.Println(content)
}
