package rabbitmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func RabbitProducer(message string) *amqp.Channel {

	_, channel := InitRabbitConnection()

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
			Body:        []byte(message),
		},
	)

	if errProduce != nil {
		fmt.Println("Cannot produce message with: ", errProduce.Error())
	}

	fmt.Println("Queue statue: ", queue)
	fmt.Println("Message produced.")
	return channel
}
