package uservalidation

import "strings"

var remainingTickets int = 50

func UserInputValidations(firstName string, lastName string, email string, ticketNumber int) (bool, bool, bool) {
	var isNameValid bool = len(firstName) > 4 && len(lastName) >= 4
	var isMailValid bool = strings.Contains(email, "@")
	var isTicketNumberValid bool = ticketNumber <= remainingTickets

	return isNameValid, isMailValid, isTicketNumberValid
}
