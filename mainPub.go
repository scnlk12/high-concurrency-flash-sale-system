package main

import (
	"strconv"
	"time"

	"github.com/scnlk12/high-concurrency-flash-sale-system/RabbitMQ"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSub("newProduct")

	for i := 0; i < 100; i++ {
		rabbitmq.PublishPub("订阅模式生产第" + strconv.Itoa(i) + "条消息")
		time.Sleep(1 * time.Second)
	}
}