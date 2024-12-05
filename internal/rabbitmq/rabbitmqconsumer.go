package rabbitmq

import (
	mysqlconnector "booking/internal/mysql"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var userInfoSlices []string

func RabbitConsumer() {

	// channel := RabbitProducer() // Be note define defer.channel.close() with error arg.
	_, channel := InitRabbitConnection()
	messages, err := channel.Consume(
		"qurywrite", // queue
		"",          // consumer
		true,        // auto ack
		false,       //exclusive
		false,       // no local
		false,       // no wait
		nil,         // args
	)

	if err != nil {
		log.Panic("Cannot consumr with error: ", err.Error())
	}

	forever := make(chan bool)
	go func() {
		for message := range messages {
			fmt.Printf("Received message: %s \n", message.Body)
			userInfos := string(message.Body[:])
			userInfoSlices := append(userInfoSlices, userInfos)
			for _, userInfo := range userInfoSlices {
				firstName := strings.Fields(userInfo)[0]
				lastName := strings.Fields(userInfo)[1]
				email := strings.Fields(userInfo)[2]
				ticketNumberString := strings.Fields(userInfo)[3]
				ticketNumber, _ := strconv.Atoi(ticketNumberString)
				mysqlconnector.Insert(firstName, lastName, email, ticketNumber)
			}
		}
	}()

	fmt.Println("Waiting for message ...")
	<-forever
}
