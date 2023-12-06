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
	harsh.Getstatus()
	harsh.NewMail()
	fmt.Printf("%+v\n", harsh.Email)

}

type user struct {
	Name         string
	Email        string
	Varification bool
	Age          int
}

func (u user) Getstatus() { // This is a Method
	fmt.Println("Is user active :", u.Varification)

} // This is creating a duplicate memory hence when asked for harsh.email it gives back original mail
func (u user) NewMail() { // we can use pointers to ensure no duplication and mail actually gets updated
	email := "testmail@gmail.com"
	fmt.Println("This user email is :", email)
}
