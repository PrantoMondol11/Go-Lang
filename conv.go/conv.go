package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	numRating, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)

	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Println("Added: ", numRating+1)
	}

}
