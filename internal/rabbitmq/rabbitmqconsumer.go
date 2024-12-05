package rabbitmq

import (
	"fmt"
	// "github.com/streadway/amqp"
	"log"
)

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
		}
	}()

	fmt.Println("Waiting for message ...")
	<-forever
}
