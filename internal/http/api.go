package main

import (
	// "os"

	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userInfoTemplate struct {
	firstName    string
	lastName     string
	email        string
	ticketNumber int
}

func unserInfoPost(body *gin.Context) {
	firstName := body.PostForm("firstnames")
	lastName := body.PostForm("lastname")
	email := body.PostForm("email")
	ticketNumberString := body.PostForm("ticketnumber")

	ticketNumber, _ := strconv.Atoi(ticketNumberString)
	userInfo := userInfoTemplate{
		firstName:    firstName,
		lastName:     lastName,
		email:        email,
		ticketNumber: ticketNumber, //ticketNumber,
	}

	fmt.Println("FirstName get from post: ", userInfo.firstName, userInfo.lastName, userInfo.email, userInfo.ticketNumber)
	//  curl -vvvv -XPOST  -d "firstnames=mamad" http://localhost/firstname
	// while true; do curl -XPOST  -d "firstnames=mamad" http://localhost/firstname && sleep 0.2 && clear ;done
}

func main() {
	server := gin.Default()
	server.GET("/hello", func(test *gin.Context) {
		test.String(http.StatusOK, "Hello world from string.")
	})

	server.POST("/firstname", unserInfoPost)

	server.GET("/", func(r *gin.Context) {
		r.JSON(200, gin.H{
			"message": "Hellow World",
		})
	})

	server.Run(":80")
}
