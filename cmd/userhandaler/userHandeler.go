package userhandaler

import (
	"fmt"
	"strconv"
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

type userData struct {
	FirstName    string
	LastName     string
	Email        string
	TicketNumber int
}

// User Info Catchup logic
func InputUserInfo() (string, string, string, int) {
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

// User Handeler Info Logic .

func UserHandeler(firstName string, lastName string, email string, ticketNumber int) (string, string, string, string) {
	firstName, lastName, email, ticketNumber = InputUserInfo()
	// Make map from user input:
	var userInputMap = make(map[string]string)
	userInputMap["firstName"] = firstName
	userInputMap["lastName"] = lastName
	userInputMap["email"] = email
	userInputMap["ticketNumber"] = strconv.Itoa(ticketNumber)

	// return firstName, lastName, email, ticketNumber
	return userInputMap["firstName"], userInputMap["lastName"], userInputMap["email"], userInputMap["ticketNumber"]
}

func UserHandelerMap(firstName string, lastName string, email string, ticketNumber int) map[string]string {
	// Make map from user input:
	var userInputMap = make(map[string]string)
	userInputMap["firstName"] = firstName
	userInputMap["lastName"] = lastName
	userInputMap["email"] = email
	userInputMap["ticketNumber"] = strconv.Itoa(ticketNumber)

	return userInputMap
}

func UserHandelerStruct(firstName string, lastName string, email string, ticketNumber int) userData {
	// Make map from user input:
	firstName, lastName, email, ticketNumber = InputUserInfo()

	var userInput = userData{
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		TicketNumber: ticketNumber,
	}

	return userInput
}
