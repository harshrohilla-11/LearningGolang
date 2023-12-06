package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")

	var fruitlist = []string{"mango", "apple", "tomato"}
	fmt.Println(fruitlist)

	fmt.Printf("fruitlist type -> %T\n", fruitlist)

	fruitlist = append(fruitlist, "peach", "banana")
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[1:3])
	fmt.Println(fruitlist)

	highscores := make([]int, 4)
	highscores[0] = 123
	highscores[1] = 987
	highscores[2] = 456
	highscores[3] = 654

	highscores = append(highscores, 753, 357, 951, 159)
	fmt.Println(highscores)

	sort.IntSlice.Sort(highscores) //sorting function
	fmt.Println(highscores)

	sort.Ints(highscores) //sorting
	fmt.Println(highscores)

	fmt.Println(sort.IntsAreSorted(highscores)) // Bool checks if slice is sorted

	//How to remove a value from slices based on index

	var courses = []string{"python", "swift", "java", "ruby", "Golang"}

	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
