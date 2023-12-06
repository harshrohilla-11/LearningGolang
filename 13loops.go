package main

import "fmt"

func main() {

	fmt.Println("Loops in Golang")

	days := []string{"Monday", "tuesday", "wednesday", "Friday", "sunday"}

	// appends string saturday at index 4
	days = append(days[:4], "saturday", days[4])
	//days = append(days[:4], append([]string{"saturday"}, days[4:]...)...) works same as above line of code

	fmt.Println(days)

	// d represents index value**Comment
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(d)
	// }

	// // d represents index value , but here we pass it in days (slice)**comment
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// Another Method**comment
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for i, day := range days {
	// 	fmt.Printf("At index value : %v days is  %v\n", i, day)
	// }

	freevalue := 0

	for freevalue < 11 { // continue , break and Goto statements

		fmt.Println(freevalue)
		freevalue++
		if freevalue == 5 {
			freevalue++
			continue

		}

		// if freevalue == 8 {
		// 	goto jump
		// }

		// if freevalue == 5 {
		// 	break
		// }
	}

	// jump:
	// 	fmt.Println("Jumped HEre")

}
