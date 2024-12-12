package httppost

import (
	// "os"

	"booking/cmd/userhandaler"
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

func UserInfoPost(body *gin.Context) {
	firstName := body.PostForm("firstnames")
	lastName := body.PostForm("lastname")
	email := body.PostForm("email")
	ticketNumberString := body.PostForm("ticketnumber")

	ticketNumber, _ := strconv.Atoi(ticketNumberString)
	UserInfo := userInfoTemplate{
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		TicketNumber: ticketNumber, //ticketNumber,
	}

	fmt.Println("Debug: From http, api --> all datas: ", UserInfo.FirstName, UserInfo.LastName, UserInfo.Email, UserInfo.TicketNumber)
	userhandaler.UserInputFromAPI(UserInfo.FirstName, UserInfo.LastName, UserInfo.Email, UserInfo.TicketNumber)
	//  curl -vvvv -XPOST  -d "firstnames=mamad" http://localhost/firstname
	// while true; do curl -XPOST  -d "firstnames=mamad" http://localhost/firstname && sleep 0.2 && clear ;done
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
