package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating:")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks", input)

}
