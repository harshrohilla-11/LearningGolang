package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome")
	fmt.Println("pls rate our pizza b/w 1 to 5 :")

	//a var that accepts time as input from user

	reader := bufio.NewReader(os.Stdin)

	//display of input

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ,", input)

	// converting accepted result string into float

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("Added 1 to your Rating :", numRating+1)
	}

}
