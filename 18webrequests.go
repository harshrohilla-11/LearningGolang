package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Welcome to WebRequets/Service in Go")

	const url = "https://lco.dev"

	Response, err := http.Get(url)
	if err != nil {
		panic(err)

	}
	fmt.Printf("The response is of type:%T", Response)
	defer Response.Body.Close() // callers responsibility to close the request

	database, err := ioutil.ReadAll(Response.Body)
	if err != nil {
		panic(err)

	}
	fmt.Println(string(database))

}
