package main

import "github.com/scnlk12/high-concurrency-flash-sale-system/RabbitMQ"

func main() {
	rabbitmq2 := RabbitMQ.NewRabbitMQTopic("newProductTopic", "gf_test.*.two")
	rabbitmq2.ReceiveTopic()
}