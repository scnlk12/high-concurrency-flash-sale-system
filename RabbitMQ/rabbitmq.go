package RabbitMQ

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// Simple 模式

// 定义 rabbitmq 连接信息 amqp://账号:密码@ip:port/vhost
const MQURL = "amqp://fan:FannaF@127.0.0.1:5672/gf"

// 创建RabbitMQ类
type RabbitMQ struct {
	conn *amqp.Connection
	channel *amqp.Channel
	// 队列名称
	QueueName string
	// 交换机
	Exchange string
	/// key
	Key string
	// 连接信息
	Mqurl string
}

// 创建RabbitMQ结构体实例
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{
		QueueName: queueName,
		Exchange: exchange,
		Key: key,
		Mqurl: MQURL,
	}

	var err error
	// 创建连接
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "创建连接错误!")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "获取channel失败!")
	return rabbitmq
}

// 断开 channel 和 connection
func (r *RabbitMQ) Destroy() {
	r.channel.Close()
	r.conn.Close()
}

// 错误处理
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

// Simple 模式 根据 QueueName 进行区分
// 创建简单模式下 RabbitMQ 实例
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	// Simple模式下，exchange 和 key 传空值，使用default配置
	rabbitmq := NewRabbitMQ(queueName, "", "")
	return rabbitmq
}

// Simple 模式下的生产者
func (r *RabbitMQ) PublishSimple(message string) {
	// 1. 申请队列 如果队列不存在会自动创建 如果存在则跳过创建
	// 保证队列存在 消息能发送到队列中
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		// 是否持久化
		false,
		// 是否自动删除
		false,
		// 是否具有排他性
		false,
		// 是否阻塞
		false,
		// 额外参数
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	// 2. 发送消息到队列中
	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		// 如果为true，会根据exchange类型和routekey规则，如果无法找到符合条件的队列会把消息返回给发送者
		false,
		// 如果为true，当exchange发送消息到队列后发现队列上没有绑定消费者，则会把消息返回给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(message),
		},
	)
}

func (r *RabbitMQ) ConsumeSimple() {
	// 1. 申请队列
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	// 2. 接收消息
	// msg 本身是一个channel
	msg, err := r.channel.Consume(
		r.QueueName,
		// 用来区分多个消费者
		"",
		// 是否自动应答
		true,
		// 是否具有排他性
		false,
		// 如果设置为true，表示不能将同一个connection中发送的消息传递给这个connection中的消费者
		false,
		// 队列消费是否阻塞
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	// 3. 消费消息
	forever := make(chan bool)

	// 启用协程处理消息
	go func() {
		for d := range msg {
			// 逻辑函数
			log.Printf("Received a message: %s", d.Body)
			fmt.Println(d.Body)
		}
	}()

	log.Printf("[*] Waiting for messages, To exit press CTRL + C")
	// 阻塞主线程，让 goroutine 持续运行
	<-forever
}