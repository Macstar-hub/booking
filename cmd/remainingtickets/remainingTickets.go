package remainingtickets

import (
	mysqlconnector "booking/internal/mysql"
)

type Tabelinfo struct {
	TicketNumber int `json:"ticketnumber"`
}

func AvailableTickets(remainingTickets int, ticketNumber int) int {

	db := mysqlconnector.MakeConnectionToDB()
	selectQuery, err := db.Query("select ticketnumber from users") // For example: db.Query("select * from users")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	for selectQuery.Next() {
		var tag Tabelinfo

		err = selectQuery.Scan(&tag.TicketNumber)
		if err != nil {
			panic(err.Error())
		}

		ticketNumber = tag.TicketNumber + ticketNumber
		defer db.Close()
	}
	remainingTickets = remainingTickets - ticketNumber
	return remainingTickets
}
