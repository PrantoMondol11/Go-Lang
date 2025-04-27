package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://chatgpt.com/c/680e340e-1aa0-8001-9568-ed84776d3cf0"

func main() {

	fmt.Println("Handling url")
	result, _ := url.Parse(myurl)
	fmt.Println(result.Path)
	fmt.Println(result.Scheme)

	qparams := result.Query()
	for _, val := range qparams {
		fmt.Println(val)
	}

	uurl := &url.URL{
		Scheme:  "https",
		Host:    "www.google.com",
		Path:    "/",
		RawPath: "user=pranto",
	}

	aurl := uurl.String()
	fmt.Println(aurl)
}
