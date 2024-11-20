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
var allAvailableTickets int

func inputUserInfo() (string, string, string, int) {
	fmt.Printf("Welcome to %v with available tickets: %v \n", confrenceName, remainingTickets)
	fmt.Println("Please Enter Your First Name: ")
	fmt.Scan(&firstName)
	fmt.Println("Please Enter Your Last Name: ")
	fmt.Scan(&lastName)
	fmt.Println("Please Enter Your Email: ")
	fmt.Scan(&email)
	fmt.Println("Please Enter Your Tickets For Bokking: ")
	fmt.Scan(&ticketNumber)

	return firstName, lastName, email, ticketNumber
}

func userHandeler(firstName string, lastName string, email string, ticketNumber int) (string, string, string, int) {
	userFirstName, userLastName, userEmail, userTicketNumber := inputUserInfo()
	fmt.Printf("User %v %v with email %v booked %v tickets \n", userFirstName, userLastName, userEmail, userTicketNumber)
	return firstName, lastName, email, ticketNumber
}

func makeArray(firstName string, lastName string) []string {
	booking := make([]string, 1)
	booking[0] = firstName + " " + lastName
	return booking
}

func availableTickets(remainingTickets int, ticketNumber int) int {
	allAvailableTickets = remainingTickets - ticketNumber
	return allAvailableTickets
}

func main() {
	fmt.Printf("Welcome to %v with available tickets: %v \n", confrenceName, remainingTickets)
	userHandeler(firstName, lastName, email, ticketNumber)
	fmt.Printf("All booking users: %v \n", makeArray(firstName, lastName))
	fmt.Printf("Total remaining tickets are %v \n", availableTickets(remainingTickets, ticketNumber))

}
