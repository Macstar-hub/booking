package main

import "fmt"

var slices []string
var firstNamess string
var lastNamess string

func makeSlicesd(firstNamess string, lastNamess string) []string {
	slices = append(slices, firstNamess+" "+lastNamess)
	return slices
}

func main2() {
	for {
		fmt.Println("Please Enter Your First Name: ")
		fmt.Scan(&firstNamess)
		fmt.Println("Please Enter Your Last Name: ")
		fmt.Scan(&lastNamess)
		fmt.Println(makeSlicesd(firstNamess, lastNamess))

	}
}
