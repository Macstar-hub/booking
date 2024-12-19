package main

import (
	"booking/cmd/remainingtickets"
	httppost "booking/internal/http"
	"booking/internal/rabbitmq"
	"fmt"

	"github.com/gin-gonic/gin"
)

var remainingTickets int = 50
var confrenceName string = "Go Lang"
var ticketNumber int

func main() {
	fmt.Printf("Welcome to %v confrence with available tickets: %v \n", confrenceName, remainingtickets.AvailableTickets(remainingTickets, ticketNumber))
	go rabbitmq.RabbitConsumer()
	server := gin.Default()
	server.LoadHTMLGlob("./web/**/*")
	server.POST("/userinfos", httppost.UserInfoPost)
	server.StaticFile("/", "./web/submit/submit.html")
	server.GET("/api/v1/getusers", httppost.BookedUsers)
	server.Run(":80")
}
