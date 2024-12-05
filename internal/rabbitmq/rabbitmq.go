package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func InitRabbitConnection() *amqp.Connection {
	fmt.Print("Debug: From main init connection")
	connection, err := amqp.Dial("amqp://user:bitnami@localhost:5672/")
	if err != nil {
		log.Panicln("Init connection has error with: ", err)
	}
	return connection
}

func RabbitProducer() *amqp.Channel {
	connectionProducer := InitRabbitConnection()
	channel, err := connectionProducer.Channel()
	if err != nil {
		log.Panic("Cannot init channel with error: ", err)
	}
	// defer channel.Close() // Move channel defination to init function and make correct arg return channel and error.

	queue, err := channel.QueueDeclare(
		"qurywrite", // queue name
		false,       // durable
		false,       // auto delete
		false,       // exclisive
		false,       // no wait
		nil,         // args
	)
	if err != nil {
		log.Panic("Cannot make queue with: ", err)
	}
	errProduce := channel.Publish(
		"queywrite", // exchange
		"testing",   // key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("testing"),
		},
	)

	if errProduce != nil {
		fmt.Println("Cannot produce message with: ", errProduce.Error())
	}

	fmt.Println("Queue statue: ", queue)
	fmt.Println("Message produced.")
	return channel
}

func RabbitConsumer() {
	channel := RabbitProducer() // Be note define defer.channel.close() with error arg.
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
			fmt.Printf("Received message: %s", message.Body)
		}
	}()

	fmt.Println("Waiting for message ...")
	<-forever
}

func main() {
	RabbitProducer()
	RabbitConsumer()
}
