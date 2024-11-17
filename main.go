package main

import (
	"fmt"
)

const totalTickets int = 50

var confrenceName string = "Go-lang"
var remainingTickets int = 50
var firstName string
var lastName string
var email string
var ticketNumber int

func main() {
	fmt.Printf("Welcome to %v with available tickets: %v \n", confrenceName, remainingTickets)

	fmt.Println("Please Enter Your First Name: ")
	fmt.Scan(&firstName)
	fmt.Println("Please Enter Your Last Name: ")
	fmt.Scan(&lastName)
	fmt.Println("Please Enter Your Email: ")
	fmt.Scan(&email)
	fmt.Println("Please Enter Your tickets: ")
	fmt.Scan(&ticketNumber)

	userHandeler(firstName, lastName, email, ticketNumber)
	remainingTickets = totalTickets - ticketNumber
	fmt.Printf("Total remaining tickets are %v \n", remainingTickets)

}

func userHandeler(firstName string, lastName string, email string, ticketNumber int) (string, string, string, int) {
	fmt.Printf("User %v %v with email %v booked %v tickets \n", firstName, lastName, email, ticketNumber)
	return firstName, lastName, email, ticketNumber
}

// func userFetchData () (string, string, string, int) {
// 	fmt.Println("Please Enter Your First Name: ")
// 	fmt.Scan(&firstName)
// 	fmt.Println("Please Enter Your Last Name: ")
// 	fmt.Scan(&lastName)
// 	fmt.Println("Please Enter Your Email: ")
// 	fmt.Scan(&email)
// 	fmt.Println("Please Enter Your tickets: ")
// 	fmt.Scan(&ticketNumber)

// 	return firstName, lastName, email, ticketNumber
// }
