package remainingtickets

func AvailableTickets(remainingTickets int, ticketNumber int) int {

	// We Sould define a slice to store last ticket available and every iteration use last slices element ...
	remainingTickets = remainingTickets - ticketNumber
	return remainingTickets
}
