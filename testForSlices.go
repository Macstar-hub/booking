package main

import "fmt"

var slices []string
var firstNamess string
var lastNamess string

func sliced(firstNamess string, lastNamess string) []string {
	slices = append(slices, firstNamess+" "+lastNamess+",")
	return slices
}

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
