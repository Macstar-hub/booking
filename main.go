package main

import (
	"fmt"
)

const tickets int = 50

var confrenceName string = "Go-lang"
var remainingTickets int = 50
var userName string
var ticketNumber int

func main() {
	fmt.Printf("Welcome to %v with available tickets: %v \n", confrenceName, remainingTickets)

	fmt.Println("Please Enter Your name: ")
	fmt.Scan(&userName)

	fmt.Println("Please Enter Your tickets: ")
	fmt.Scan(&ticketNumber)

	userHandeler(userName, ticketNumber)

}

func userHandeler(userName string, ticketNumber int) (string, int) {
	fmt.Printf("User %v booked %v tickets \n", userName, ticketNumber)
	return userName, ticketNumber
}
