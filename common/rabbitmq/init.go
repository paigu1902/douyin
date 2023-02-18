package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
)

const MQURL = "amqp://admin:RabbitMQ@http://172.30.16.1:5672/"

type RabbitMQ struct {
	Conn       *amqp.Connection
	Channel    *amqp.Channel
	QueueName  string
	Exchange   string
	RoutingKey string
	Mqurl      string
}

var RMQ *RabbitMQ

// 初始化RabbitMQ连接与通道
func InitRabbitMQ(queueName string) *RabbitMQ {
	RMQ := RabbitMQ{
		QueueName: queueName,
		//Exchange:   exchange,
		//RoutingKey: routingKey,
		Mqurl: MQURL,
	}
	var err error
	RMQ.Conn, err = amqp.Dial(RMQ.Mqurl)
	CheckErr(err, "Establish Connection Failed")
	RMQ.Channel, err = RMQ.Conn.Channel()
	CheckErr(err, "Establish Channel Failed")
	//go RMQ.Consume()
	return &RMQ
}

// 关闭RabbitMQ连接与通道
func (mq *RabbitMQ) CloseRabbitMQ() {
	mq.Conn.Close()
	mq.Channel.Close()
}

// 检查连接错误
func CheckErr(err error, meg string) {
	if err != nil {
		log.Fatalf("%s:%s\n", meg, err)
	}
}
