package main

import (
	"fmt"
)

var firstName string
var lastName string
var bookingSlice []string

func main() {
	for {
		fmt.Println("Please Enter Your First Name: ")
		fmt.Scan(&firstName)
		fmt.Println("Please Enter Your Last Name: ")
		fmt.Scan(&lastName)
		bookingSlice := append(bookingSlice, firstName+" "+lastName)
		fmt.Println(bookingSlice)

	}
}
