package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Web server")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:3000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)
	var responseString strings.Builder

	content, err := ioutil.ReadAll(response.Body)
	bytCount, _ := responseString.Write(content)
	fmt.Println(bytCount)
	fmt.Println(responseString.String())
}
