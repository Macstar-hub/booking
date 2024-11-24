package main

import (
	"fmt"
	"strings"
)

var slices []string
var firstNamess string
var lastNamess string
var firstNameSlice []string
var name string

func makeSlicesd(firstNamess string, lastNamess string) []string {
	var testFirstName []string
	slices = append(slices, firstNamess+" "+lastNamess)
	for i, firstNameSlice := range slices {
		fmt.Println(i)
		name = strings.Fields(firstNameSlice)[0]
		testFirstName = append(testFirstName, name)
		fmt.Println("All Names are: ", testFirstName)
	}
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
