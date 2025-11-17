package main

import "github.com/scnlk12/high-concurrency-flash-sale-system/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSub("newProduct")
	rabbitmq.ReceiveSub()
}