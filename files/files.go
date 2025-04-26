package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Hello in files")
	content := "This need to go in a file - Pranto"
	file, err := os.Create("./go.txt")
	if err != nil {
		panic(err)
	}
	len, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length is :", len)
	file.Close()
	readfile("./go.txt")
}

func readfile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("Text data inside is \n:", string(databyte))

}
