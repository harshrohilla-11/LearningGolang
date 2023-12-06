package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&payemntid=ghb5465ghb"

func main() {
	fmt.Println("Handling URL's in GOlang")
	fmt.Println(myurl) // getting Url

	//parsing
	result, _ := url.Parse(myurl)
	fmt.Println(result)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())// port is actually a method
	fmt.Println(result.RawQuery)

	queryparameters := result.Query()
	fmt.Printf("Type :%T\n", queryparameters)

	fmt.Println(queryparameters["coursename"]) // as these are key value pairs , can access them using key
	fmt.Println(queryparameters["payemntid"])

	for _, val := range queryparameters { // iterating through parameters
		fmt.Println("Parameters are :", val)

	}

	partsofurl := &url.URL{ //NEVER FORGET &[Pointer] HERE
		Scheme:   "https",
		Host:     "localMachine",
		Path:     "/login",
		RawQuery: "user=Harsh",
	}
	//This is url that we are constructing using above &url.URl mathod
	anotherurl := partsofurl.String()
	fmt.Println(anotherurl)
}
