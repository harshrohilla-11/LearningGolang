package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in GO")

	content := "This is a test file for Go - Harsh"

	file, err := os.Create("./testfile.txt")
	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Println("The File length is :", length)
	FileReader("./testfile.txt")

	defer file.Close()
}
func FileReader(filname string) {
	database, err := ioutil.ReadFile(filname)
	checkNilError(err)
	fmt.Println("The data in this file testfile.txt is:\n", string(database))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}

}
