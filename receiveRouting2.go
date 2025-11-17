package main

import "github.com/scnlk12/high-concurrency-flash-sale-system/RabbitMQ"

func main() {
	rabbitmq2 := RabbitMQ.NewRabbitMQRouting("newProductRouting", "gf_test_2")
	rabbitmq2.ReceiveRouting()
}