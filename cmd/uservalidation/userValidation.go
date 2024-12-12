package uservalidation

import (
	"booking/cmd/remainingtickets"
	"strings"
)

var remainingTickets int = 50

func UserInputValidations(firstName string, lastName string, email string, ticketNumber int) (bool, bool, bool) {
	var isNameValid bool = len(firstName) > 4 && len(lastName) >= 4
	var isMailValid bool = strings.Contains(email, "@")
	remainingticketSql := remainingtickets.AvailableTickets(remainingTickets, ticketNumber)
	var isTicketNumberValid bool = ticketNumber <= remainingticketSql

	return isNameValid, isMailValid, isTicketNumberValid
}
