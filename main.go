package main

import (
	"booking/remainingtickets"
	"fmt"
	"strings"
)

var totalTicket int = 50
var remainingTickets int = 50
var booking []string
var confrenceName string = "Go-lang"
var email string
var ticketNumber int
var firstName string
var lastName string
var slicess []string
var allFirstName []string
var names string
var reservedTicket int

// var allAvailableTickets int

func inputUserInfo() (string, string, string, int) {
	fmt.Println("Please Enter Your First Name: ")
	fmt.Scan(&firstName)
	fmt.Println("Please Enter Your Last Name: ")
	fmt.Scan(&lastName)
	fmt.Println("Please Enter Your Email: ")
	fmt.Scan(&email)
	fmt.Println("Please Enter Your Tickets For Booking: ")
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

func makeSlicesd(firstNamess string, lastNamess string) ([]string, []string) {
	var testFirstName []string
	slicess = append(slicess, firstNamess+" "+lastNamess)
	for _, firstNameSlice := range slicess {
		names = strings.Fields(firstNameSlice)[0]
		testFirstName = append(testFirstName, names)
	}
	return slicess, testFirstName
}

func main() {
	fmt.Printf("Welcome to %v with available tickets: %v \n", confrenceName, remainingTickets)

	for {
		fmt.Printf("var remainingTickets in first func: %v \n", remainingTickets)
		userHandeler(firstName, lastName, email, ticketNumber)
		allUsers, allFirsnames := makeSlicesd(firstName, lastName)
		fmt.Printf("All users are: %v And all just firstnames: %v\n", allUsers, allFirsnames)
		remainingTickets = remainingtickets.AvailableTickets(remainingTickets, ticketNumber)

		notEnoughTickets := remainingTickets <= 0
		if notEnoughTickets {
			fmt.Printf("All Ticket Remaining: %v \n", remainingTickets)
			break
		} else {
			fmt.Println("You Can Continue with total remaning tickets: %v \n", remainingTickets)
		}
	}
}
