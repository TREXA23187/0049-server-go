package console_api

import (
	"0049-server-go/global"
	"0049-server-go/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"log"
)

type TaskOperateRequest struct {
	TaskId    string `json:"task_id" binding:"required" msg:"Please enter task id"`
	Operation string `json:"operation" binding:"required" msg:"Please enter operation"`
}

func (ConsoleApi) TaskOperateView(ctx *gin.Context) {
	var cr TaskOperateRequest
	if err := ctx.ShouldBindJSON(&cr); err != nil {
		res.FailWithValidMsg(err, &cr, ctx)
		return
	}

	ch, err := global.MQ.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"task_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Failed to declare a queue", err)
	}

	body := cr.TaskId

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	if err != nil {
		log.Fatal("Failed to publish a message", err)
	}

	res.OkWithMessage(fmt.Sprintf("%s the task successfully", string(cr.Operation)), ctx)
}
