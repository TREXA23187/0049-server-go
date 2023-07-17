package configs

import (
	"bytes"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

type MQ struct {
	IP       string `json:"ip" yaml:"ip"`
	Port     int    `json:"port" yaml:"port"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
}

func (mq MQ) Addr() string {
	return fmt.Sprintf("%s:%d", mq.IP, mq.Port)
}

func (mq MQ) URL() string {
	return fmt.Sprintf("amqp://%s:%s@%s/", mq.Username, mq.Password, mq.Addr())
}

func Receive() {
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	defer conn.Close()
	ch, _ := conn.Channel()
	defer ch.Close()
	q, _ := ch.QueueDeclare(
		"task_finished_queue", // name
		true,                  // durable
		false,                 // delete when unused
		false,                 // exclusive
		false,                 // no-wait
		nil,                   // arguments
	)

	ch.Qos(
		1,     // prefetch count 服务器将在收到确认之前将那么多消息传递给消费者。
		0,     // prefetch size  服务器将尝试在收到消费者的确认之前至少将那么多字节的交付保持刷新到网络
		false, // 当 global 为 true 时，这些 Qos 设置适用于同一连接上所有通道上的所有现有和未来消费者。当为 false 时，Channel.Qos 设置将应用于此频道上的所有现有和未来消费者
	)

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

			// func Count(s, sep [] byte ) int  计算s中sep的非重叠实例数。如果sep是空切片，则Count返回1+s中UTF-8编码的代码点数
			dotCount := bytes.Count(d.Body, []byte(".")) // 返回 . 的个数

			t := time.Duration(dotCount) // 表示为 int64 纳秒计数
			time.Sleep(t * time.Second)  //将当前 goroutine 暂停至少持续时间d

			log.Printf("Done")
			d.Ack(false)

		}
	}()

	<-forever
}
