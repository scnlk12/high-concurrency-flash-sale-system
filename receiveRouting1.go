package main

import "github.com/scnlk12/high-concurrency-flash-sale-system/RabbitMQ"

func main() {
	rabbitmq1 := RabbitMQ.NewRabbitMQRouting("newProductRouting", "gf_test_1")
	rabbitmq1.ReceiveRouting()
}