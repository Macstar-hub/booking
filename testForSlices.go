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
		fmt.Println("Enter words: ")
		fmt.Scan(&firstNamess)
		fmt.Scan(&lastNamess)
		fmt.Println(sliced(firstNamess, lastNamess))
	}
}
