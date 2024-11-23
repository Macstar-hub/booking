package main

import (
	"fmt"
	"strings"
)

var totalTicket int = 50
var bookingSlice []string
var booking []string
var confrenceName string = "Go-lang"
var firstName string
var lastName string
var email string
var ticketNumber int
var remainingTickets int

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

func makeSlice(firstName string, lastName string) []string {
	bookingSlice := append(bookingSlice, firstName+" "+lastName)

	return bookingSlice
}

func firstNames(bookingSlice []string) []string {
	var booking = makeSlice(firstName, lastName)
	var firstNames []string
	for _, bookingFirstName := range booking {
		var name = strings.Fields(bookingFirstName)
		firstNames = append(firstNames, name[0])
	}

	return firstNames
}

func availableTickets(ticketNumber int) int {
	var confCapacity []int
	totalAvailableTickets := append(confCapacity, totalTicket)
	totalTicket = totalAvailableTickets[len(totalAvailableTickets)-1] - ticketNumber
	return totalTicket
}

func main2() {
	fmt.Printf("Welcome to %v with available tickets: %v \n", confrenceName, remainingTickets)

	for {
		userHandeler(firstName, lastName, email, ticketNumber)
		bookingSlice := append(booking, firstName+" "+lastName)

		fmt.Printf("All booking users with arrays: %v \n", makeArray(firstName, lastName))
		fmt.Println("======================================================================================================")
		fmt.Printf("All booking users with all booking slices: %v \n", bookingSlice)
		fmt.Printf("Total remaining tickets are %v \n", availableTickets(ticketNumber))
		// fmt.Printf("All first name are: %v \n", firstNames(allbooking))

	}
}
