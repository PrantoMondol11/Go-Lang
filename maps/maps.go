package main

import "fmt"

func main() {

	languages := make(map[string]string)
	languages["js"] = "Javascript"
	languages["Go"] = "go"
	languages["py"] = "python"
	languages["rb"] = "ruby"

	fmt.Println("js Shows:", languages["js"])
	delete(languages, "rb")
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v \n", key, value)
	}
}
