package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/scnlk12/high-concurrency-flash-sale-system/RabbitMQ"
)

func main() {
	rabbitmq1 := RabbitMQ.NewRabbitMQRouting("newProductRouting", "gf_test_1")
	rabbitmq2 := RabbitMQ.NewRabbitMQRouting("newProductRouting", "gf_test_2")

	for i := 1; i < 100; i++ {
		rabbitmq1.PublishRouting("路由模式rb1生产第" + strconv.Itoa(i) + "条消息")
		rabbitmq2.PublishRouting("路由模式rb2生产第" + strconv.Itoa(i) + "条消息")
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}