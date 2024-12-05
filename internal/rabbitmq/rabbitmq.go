package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

func InitRabbitConnection() (*amqp.Connection, *amqp.Channel) {
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
