package core

import (
	"0049-server-go/global"
	"github.com/streadway/amqp"
	"log"
)

var rabbitConn *amqp.Connection

func ConnectRabbitMQ() *amqp.Connection {
	mqConf := global.Config.MQ

	var err error
	rabbitConn, err = amqp.Dial(mqConf.URL())
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ", err)
	}

	return rabbitConn
}
