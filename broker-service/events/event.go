package events

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		"logs_topic", //name
		"topic",      //
		true,         //durable
		false,        //autho-deleted?
		false,        //internal?
		false,        //no wait?
		nil,          //arguments?
	)
}

func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",    // name
		false, //durable
		false, //delete when unused
		true,  //eclusice
		false, //no-wait
		nil,   //arguments
	)
}
