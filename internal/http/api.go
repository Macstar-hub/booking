package httppost

import (
	// "os"

	"booking/cmd/remainingtickets"
	"booking/cmd/userhandaler"
	"booking/cmd/uservalidation"
	"booking/internal/rabbitmq"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// test
type userInfoTemplate struct {
	FirstName    string
	LastName     string
	Email        string
	TicketNumber int
}

var remainingTickets int = 50

func UserInfoPost(body *gin.Context) {
	firstName := body.PostForm("firstnames")
	lastName := body.PostForm("lastnames")
	email := body.PostForm("emails")
	ticketNumberString := body.PostForm("ticketnumbers")

	ticketNumber, _ := strconv.Atoi(ticketNumberString)
	UserInfo := userInfoTemplate{
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		TicketNumber: ticketNumber,
	}
	userData := userhandaler.UserHandelerStruct(UserInfo.FirstName, UserInfo.LastName, UserInfo.Email, UserInfo.TicketNumber)
	firstName, lastName, email, ticketNumber = userData.FirstName, userData.LastName, userData.Email, userData.TicketNumber
	allUsers, allFirsnames := userhandaler.MakeSlicesd(firstName, lastName, email, ticketNumber)
	fmt.Printf("All users are: %v And all just firstnames: %v\n", allUsers, allFirsnames)

	isNameValid, isMailValid, isTicketNumberValid := uservalidation.UserInputValidations(firstName, lastName, email, ticketNumber)
	if isNameValid && isMailValid && isTicketNumberValid {
		remainingTickets = remainingtickets.AvailableTickets(remainingTickets, ticketNumber)

		// Produce test info to rabbitmq
		rabbitmq.RabbitProducer(firstName, lastName, email, ticketNumber)
		go rabbitmq.RabbitConsumer()
		//

		notEnoughTickets := remainingTickets <= 0
		if notEnoughTickets {
			fmt.Printf("Booking Failed With Ticket Remaining: %v \n", remainingTickets)
			// break
		} else {
			fmt.Printf("User %v %v With Email %v Booked Successfuly %v Tickets \n", firstName, lastName, email, ticketNumber)
		}

		// Make user feedback guid
	} else {
		if !isNameValid {
			fmt.Println("Please Enter Correct FirstName And LastName ... ")
			fmt.Printf("FisrtName that enterd: %v And lastName that enterd: %v", firstName, lastName)
			// break
		}
		if !isMailValid {
			fmt.Println("Please Enter Correct Email Address ... ")
			// break
		}
		if !isTicketNumberValid {
			fmt.Printf("Please Select Ticket Number In Range Remaining Tickets: %v \n", remainingTickets)
			// break
		}
	}

	userhandaler.UserInputFromAPI(UserInfo.FirstName, UserInfo.LastName, UserInfo.Email, UserInfo.TicketNumber)

}

// func ServerRun() {
// 	server := gin.Default()
// 	server.GET("/hello", func(test *gin.Context) {
// 		test.String(http.StatusOK, "Hello world from string.")
// 	})

// 	server.POST("/userinfos", UserInfoPost)

// 	server.GET("/", func(r *gin.Context) {
// 		r.JSON(200, gin.H{
// 			"message": "Hellow World",
// 		})
// 	})

// 	server.Run(":80")
// }
