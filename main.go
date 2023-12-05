package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Welcome to Time and date tutorial in Golang ")

	time_now := time.Now() // present time

	// formatting time look
	fmt.Println(time_now.Format("01-02-2006 monday"))
	instant := time.Date(2002, time.January, 01, 2, 12, 00, 00, time.Local)

	// formatting time look
	fmt.Println(instant.Format("01-02-2006 15:04:05 monday"))

}
