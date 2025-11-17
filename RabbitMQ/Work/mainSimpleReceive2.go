package main

import "github.com/scnlk12/high-concurrency-flash-sale-system/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("gf_test")
	rabbitmq.ConsumeSimple()
}