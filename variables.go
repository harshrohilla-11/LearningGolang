package main

import "fmt"

//logout := "u have logged out"  not allowed on global

const first = "hiiiii"

const second = " byeee" // capital P makes it Public

func main() {
	var isloggedin bool = true
	fmt.Println(isloggedin)
	fmt.Printf("this is the type : %T \n", isloggedin)

	var smallvalue uint8 = 255
	fmt.Println(smallvalue)
	fmt.Printf("this is the type : %T \n", smallvalue)

	var newsmallvalue int = 257
	fmt.Println(newsmallvalue)
	fmt.Printf("this is the type : %T \n", newsmallvalue)

	var smallfloat float64 = 456.646949649694
	fmt.Println(smallfloat)
	fmt.Printf("this is the type : %T \n", smallfloat)

	//defalut values and some allias
	var anothervariable string
	fmt.Println(anothervariable)
	fmt.Printf("this is the type : %T \n", anothervariable)

	loginpage := "hello sir , pls log in "
	fmt.Println(loginpage)

	fmt.Println(first)
	fmt.Printf("this is the type : %T \n", first)

	fmt.Println(second)
	fmt.Printf("this is the type : %T \n", second)

}
