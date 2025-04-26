package main

import "fmt"

func main() {
	fmt.Println("Hello")
	defer fmt.Println("Hello word")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Pranto")
	mydefer()
}

func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
