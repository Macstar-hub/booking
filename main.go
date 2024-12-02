package main

import (
	"booking/cmd/remainingtickets"
	"booking/cmd/userhandaler"
	"fmt"
	"strings"
)

var totalTicket int = 50
var remainingTickets int = 50
var booking = make([]map[string]string, 0)
var confrenceName string = "Go-lang"
var email string

var ticketNumber int
var firstName string
var lastName string
var slicess []string
var allFirstName []string
var names string
var reservedTicket int
var bookingMap map[string]string

func makeSlicesd(firstNamess string, lastNamess string, email string, ticketNumber int) ([]string, []string) {
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
		// Force to user struct function.
		userData := userhandaler.UserHandelerStruct(firstName, lastName, email, ticketNumber)
		firstName, lastName, email, ticketNumber = userData.FirstName, userData.LastName, userData.Email, userData.TicketNumber
		fmt.Println("Debug: inside main -> for.")
		allUsers, allFirsnames := makeSlicesd(firstName, lastName, email, ticketNumber)
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
				fmt.Printf("FisrtName that enterd: %v And lastName that enterd: %v", firstName, lastName)
				break
			}
			if !isMailValid {
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
