package main

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

func main() {

	// Declare exchange and queue
	conn, err := initConnection()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	exchangeName := "payment-exchange"
	dlxExchangeName := "dlx-payment-exchange"
	queueName := "payment-queue"
	dlxQueueName := "dlx-payment-queue"
	err = declareExchangeAndQueueWithDlx(ch, exchangeName, dlxExchangeName, queueName)
	if err != nil {
		panic(err)
	}

	err = declareExchangeAndQueue(ch, dlxExchangeName, dlxQueueName, true)
	if err != nil {
		panic(err)
	}

	// Publish message
	sendMsg(ch, exchangeName, "Hello World", 2000)

	// // Consume message
	// msgs, err := consumeMsg(ch, queueName)
	// if err != nil {
	// 	panic(err)
	// }

	// for d := range msgs {
	// 	println("this is normal queue ", string(d.Body))
	// 	d.Nack(false, false)
	// }

	// check dlx queue
	msgs, err := consumeMsg(ch, dlxQueueName)
	if err != nil {
		panic(err)
	}

	for d := range msgs {
		println("this is dlx queue", string(d.Body))
		// d.Nack(false, false)
	}

}

func initConnection() (*amqp.Connection, error) {
	uri := "amqp://rabbitmq:rabbitmq@localhost:5672/"

	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func sendMsg(ch *amqp.Channel, exchangeName string, body interface{}, delay int) error {
	jsPayload, err := json.Marshal(body)
	if err != nil {
		return err
	}
	header := make(amqp.Table)
	header["x-delay"] = delay
	err = ch.Publish(exchangeName, "", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(jsPayload),
		Headers:     header,
	})
	if err != nil {
		return err
	}

	return nil
}

func consumeMsg(ch *amqp.Channel, queueName string) (<-chan amqp.Delivery, error) {
	msgs, err := ch.Consume(queueName, "", false, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}

func declareExchangeAndQueueWithDlx(ch *amqp.Channel, exchangeName string, dlxExchangeName string, queueName string) error {
	// Declare an exchange
	err := ch.ExchangeDeclare(exchangeName, amqp.ExchangeTopic, true, false, false, false, nil)
	if err != nil {
		return err
	}

	// Declare a queue
	args := make(amqp.Table)
	args["x-dead-letter-exchange"] = dlxExchangeName
	_, err = ch.QueueDeclare(queueName, true, false, false, false, args)
	if err != nil {
		return err
	}

	// create binding between exchange and queue
	err = ch.QueueBind(queueName, "#", exchangeName, false, nil)
	if err != nil {
		return err
	}

	return nil
}

func declareExchangeAndQueue(ch *amqp.Channel, exchangeName string, queueName string, isDelayedExchange bool) error {
	// Declare an exchange
	if isDelayedExchange {
		argsEx := make(amqp.Table)
		argsEx["x-delayed-type"] = "direct"
	}
	err := ch.ExchangeDeclare(exchangeName, amqp.ExchangeTopic, true, false, false, false, nil)
	if err != nil {
		return err
	}

	// Declare a queue
	args := make(amqp.Table)
	_, err = ch.QueueDeclare(queueName, true, false, false, false, args)
	if err != nil {
		return err
	}

	// create binding between exchange and queue
	err = ch.QueueBind(queueName, "#", exchangeName, false, nil)
	if err != nil {
		return err
	}

	return nil
}
