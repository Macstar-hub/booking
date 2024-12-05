package rabbitmq

import (
	"fmt"
	"log"

	"strconv"

	"github.com/streadway/amqp"
)

type Message struct {
	FirstName    string
	Lastname     string
	Email        string
	TicketNumber int
}

func RabbitProducer(firstName string, lastName string, email string, ticketNumber int) *amqp.Channel {

	_, channel := InitRabbitConnection()

	// message := Message{
	// 	FirstName:    firstName,
	// 	Lastname:     lastName,
	// 	Email:        email,
	// 	TicketNumber: ticketNumber,
	// }
	message := fmt.Sprintf(firstName + " " + lastName + " " + email + " " + strconv.Itoa(ticketNumber))

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
