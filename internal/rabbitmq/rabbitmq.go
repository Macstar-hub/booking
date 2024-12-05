package rabbitmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func InitRabbitConnection() (*amqp.Connection, *amqp.Channel) {
	fmt.Print("Debug: From main init connection")
	connection, err := amqp.Dial("amqp://user:bitnami@localhost:5672/")
	if err != nil {
		log.Panicln("Init connection has error with: ", err)
	}
	channel, err := connection.Channel()
	if err != nil {
		log.Panic("Cannot init channel with error: ", err)
	}

	return connection, channel
}
