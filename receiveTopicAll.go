package main

import "github.com/scnlk12/high-concurrency-flash-sale-system/RabbitMQ"

func main() {
	rabbitmq1 := RabbitMQ.NewRabbitMQTopic("newProductTopic", "#")
	rabbitmq1.ReceiveTopic()
}