package main

import (
	"booking/cmd/remainingtickets"
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
	inputUserInfo()
	return firstName, lastName, email, ticketNumber
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

func userInputValidations(firstName string, lastName string, email string, ticketNumber int) (bool, bool, bool) {
	var isNameValid bool = len(firstName) > 4 && len(lastName) >= 4
	var isMailValid bool = strings.Contains(email, "@")
	var isTicketNumberValid bool = ticketNumber <= remainingTickets

	return isNameValid, isMailValid, isTicketNumberValid
}

func main() {
	fmt.Printf("Welcome to %v with available tickets: %v \n", confrenceName, remainingTickets)

	for {
		userHandeler(firstName, lastName, email, ticketNumber)
		allUsers, allFirsnames := makeSlicesd(firstName, lastName)
		fmt.Printf("All users are: %v And all just firstnames: %v\n", allUsers, allFirsnames)

		// Add user input validation
		isNameValid, isMailValid, isTicketNumberValid := userInputValidations(firstName, lastName, email, ticketNumber)
		if isNameValid && isMailValid && isTicketNumberValid {
			remainingTickets = remainingtickets.AvailableTickets(remainingTickets, ticketNumber)
			notEnoughTickets := remainingTickets <= 0
			if notEnoughTickets {
				fmt.Printf("Booking Failed With Ticket Remaining: %v \n", remainingTickets)
				break
			} else {
				fmt.Printf("User %v %v With Email %v Booked Successfuly %v Tickets \n", firstName, lastName, email, ticketNumber)
			}

			// Make user feedback guid
		} else {
			if !isNameValid {
				fmt.Println("Please Enter Correct FirstName And LastName ... ")
				break
			}
			if !isNameValid {
				fmt.Println("Please Enter Correct Email Address ... ")
				break
			}
			if !isTicketNumberValid {
				fmt.Printf("Please Select Ticket Number In Range Remaining Tickets: %v \n", remainingTickets)
				break
			}
		}
	}
}
