package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/scnlk12/high-concurrency-flash-sale-system/RabbitMQ"
)

func main() {
	rabbitmq1 := RabbitMQ.NewRabbitMQTopic("newProductTopic", "gf_test.topic.one")
	rabbitmq2 := RabbitMQ.NewRabbitMQTopic("newProductTopic", "gf_test.topic.two")

	for i := 0; i < 100; i++ {
		rabbitmq1.PublishTopic("Hello gf_test topic one!" + strconv.Itoa(i))
		rabbitmq2.PublishTopic("Hello gf_test topic two!" + strconv.Itoa(i))

		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}