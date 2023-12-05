package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "weclcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Pls rate our pizza:")

	// comma ok  , err || err

	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating :", input)
	fmt.Printf("type is %T \n", reader)
	fmt.Printf("type is %T \n", input)

	var logincounter int64 = 20

	//this is used to take integer as input

	// fmt.Println("Enter your login Count :")
	// fmt.Scanln(&logincounter)

	// fmt.Printf("Enter your value :")
	// fmt.Scanf("%d", &logincounter) --> This Too

	fmt.Println(logincounter)

}
