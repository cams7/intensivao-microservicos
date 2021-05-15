package queue

import (
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

func Connect() *amqp.Channel {
	dsn := "amqp://" + os.Getenv("RABBITMQ_DEFAULT_USER") + ":" + os.Getenv("RABBITMQ_DEFAULT_PASS") + "@" + os.Getenv("RABBITMQ_DEFAULT_HOST") + ":" + os.Getenv("RABBITMQ_DEFAULT_PORT") + os.Getenv("RABBITMQ_DEFAULT_VHOST")

	conn, err := amqp.Dial(dsn)
	if err != nil {
		panic(err.Error())
	}

	channel, err := conn.Channel()
	if err != nil {
		panic(err.Error())
	}
	return channel
}

func Notify(payload []byte, exchange string, routingKey string, ch *amqp.Channel) {

	err := ch.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(payload),
		})

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Message sent")
}

func StartConsuming(queue string, ch *amqp.Channel, in chan []byte) {

	q, err := ch.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err.Error())
	}

	msgs, err := ch.Consume(
		q.Name,
		"checkout",
		true, //Remove da fila após a leitura
		false,
		false,
		false,
		nil)

	if err != nil {
		panic(err.Error())
	}

	//Goroutines - Função assicrona do go
	go func() {
		for m := range msgs {
			//O conteudo da mensagem está sendo recebido por um channel
			//Obs.: O channel só pode receber a próxima mensagem quando for esvaziado
			in <- []byte(m.Body)
		}
		close(in)
	}()
}
