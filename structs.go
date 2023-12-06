package main

import (
	"fmt"
)

func main() {

	// alternative for classes , we do not have classes in Golang
	fmt.Println("Structs In Go")

	//no inheritance in Golanf : No parents or super
	harsh := user{"Harsh", "harsh071@gmail.com", true, 21}
	fmt.Println(harsh)
	fmt.Printf("%+v\n", harsh)
	fmt.Printf("The name is :%v\nand Age is %v\n:", harsh.Name, harsh.Age)

}

type user struct {
	Name         string
	Email        string
	Varification bool
	Age          int
}
