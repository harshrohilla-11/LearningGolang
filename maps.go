package main

import (
	"fmt"
)

func main() {

	fmt.Println(" Maps in Golang ")

	languages := make(map[string]string)
	languages["Js"] = "javascript"
	languages["Py"] = "python"
	languages["Ru"] = "ruby"
	languages["Go"] = "Golang"

	fmt.Println(languages)

	fmt.Println("Js is short for :", languages["Js"])
	//fmt.Println("Py is abb. for :", languages["python"])<<gives no output >>
	delete(languages, "Ru")
	fmt.Println(languages)

	//loops in maps

	for key, value := range languages {
		fmt.Printf("The Abv. %v stands for %v\n", key, value)

	}

}
