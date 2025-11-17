package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/scnlk12/high-concurrency-flash-sale-system/RabbitMQ"
)

func main() {
	// queueName
	rabbitmq := RabbitMQ.NewRabbitMQSimple("gf_test")

	// rabbitmq.PublishSimple("Hello World!")
	for i := 1; i < 100; i++ {
		rabbitmq.PublishSimple("Hello World!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}