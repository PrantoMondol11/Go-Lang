package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	Pranto := User{"Pranto Mondol", "mondol@gmail.com", true, 22}

	fmt.Printf("Details are : %+v\n", Pranto)
}
