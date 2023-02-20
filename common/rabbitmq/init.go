package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
	"paigu1902/douyin/common/config"
	"strings"
)

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
	builder := strings.Builder{}
	builder.WriteString("amqp://")
	builder.WriteString(config.C.RabbitMQ.Username)
	builder.WriteString(":")
	builder.WriteString(config.C.RabbitMQ.Password)
	builder.WriteString("@")
	builder.WriteString(config.C.RabbitMQ.Host)
	builder.WriteString(":")
	builder.WriteString(config.C.RabbitMQ.Port)
	builder.WriteString(config.C.RabbitMQ.Vhost)

	RMQ := RabbitMQ{
		QueueName: queueName,
		Mqurl:     builder.String(),
	}
	var err error
	RMQ.Conn, err = amqp.Dial(RMQ.Mqurl)
	CheckErr(err, "Establish Connection Failed")
	RMQ.Channel, err = RMQ.Conn.Channel()
	CheckErr(err, "Establish Channel Failed")
	return &RMQ
}

// 关闭RabbitMQ连接与通道
func (favo *RabbitMQ) CloseRabbitMQ() {
	favo.Conn.Close()
	favo.Channel.Close()
}

// 检查连接错误
func CheckErr(err error, meg string) {
	if err != nil {
		log.Fatalf("%s:%s\n", meg, err)
	}
}
