package core

import (
	"0049-server-go/global"
	"bytes"
	"github.com/streadway/amqp"
	"log"
	"time"
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

func InitRabbitMQConsuming() {
	global.Log.Infof("The rabbitmq is listening on: %s", global.Config.MQ.Addr())
	ch, _ := rabbitConn.Channel()
	defer ch.Close()

	q, _ := ch.QueueDeclare(
		"task_finished_queue", // name
		true,                  // durable
		false,                 // delete when unused
		false,                 // exclusive
		false,                 // no-wait
		nil,                   // arguments
	)

	err := ch.Qos(
		1,
		0,
		false,
	)
	if err != nil {
		log.Println("Failed to set QoS", err)
		return
	}

	messages, _ := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	forever := make(chan bool)
	go func() {
		for d := range messages {
			log.Printf("Received a message: %s", d.Body)

			dotCount := bytes.Count(d.Body, []byte("."))

			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)

			log.Printf("Done")
			err = d.Ack(false)
			if err != nil {
				log.Println("Failed to send Ack", err)
				return
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
