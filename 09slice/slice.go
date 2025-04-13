package main

import (
	"fmt"
	"sort"
)

func main() {

	var fruitList = []string{"Apple", "Tomato", "Orange"}
	fruitList = append(fruitList, "Mango")
	fmt.Println(fruitList)
	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)
	fmt.Println("Hello")

	highscores := make([]int, 4)
	highscores[0] = 234
	highscores[1] = 233
	highscores[2] = 235
	highscores[3] = 245
	highscores = append(highscores, 2321, 2334, 567, 867, 4, 3)
	sort.Ints(highscores)

	fmt.Println(highscores)

	highscores = append(highscores[:2], highscores[4:]...)
	fmt.Println(highscores)
}
