package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"paigu1902/douyin/common/models"
	"strconv"
	"strings"
)

var RmqFavoAdd *RabbitMQ
var RmqFavoDel *RabbitMQ

// 初始化RabbitMQ
func InitFavoRmq() {
	RmqFavoAdd = InitRabbitMQ("favoAdd")
	go RmqFavoAdd.Consume()
	RmqFavoDel = InitRabbitMQ("favoDel")
	go RmqFavoDel.Consume()
}

// 生产者
func (favo *RabbitMQ) Publish(msg string) {
	// 声明队列
	_, err := favo.Channel.QueueDeclare(
		favo.QueueName,
		true,  // 是否持久化
		false, // 是否自动删除(前提是至少有一个消费者连接到这个队列，之后所有与这个队列连接的消费者都断开时，才会自动删除。注意：生产者客户端创建这个队列，或者没有消费者客户端与这个队列连接时，都不会自动删除这个队列)
		false, // 是否为排他队列（排他的队列仅对“首次”声明的conn可见[一个conn中的其他channel也能访问该队列]，conn结束后队列删除）
		false, // 是否阻塞
		nil,   //额外属性
	)
	if err != nil {
		//panic(err)
		log.Printf("Declare Queue Failed", err)
		return
	}
	//// 声明交换器
	//err = favo.Channel.ExchangeDeclare(
	//	favo.Exchange, // 交换器名
	//	"topic",     // exchange type：一般用fanout、direct、topic
	//	true,        // 是否持久化
	//	false,       // 是否自动删除（自动删除的前提是至少有一个队列或者交换器与这和交换器绑定，之后所有与这个交换器绑定的队列或者交换器都与此解绑）
	//	false,       // 设置是否内置的。true表示是内置的交换器，客户端程序无法直接发送消息到这个交换器中，只能通过交换器路由到交换器这种方式
	//	false,       // 是否阻塞
	//	nil,         // 额外属性
	//)
	//if err != nil {
	//	fmt.Println("Declare Exchange Failed", err)
	//	return
	//}
	//// 建立Binding(可随心所欲建立多个绑定关系)
	//err = favo.Channel.QueueBind(
	//	favo.QueueName,  // 绑定的队列名称
	//	favo.RoutingKey, // bindkey 用于消息路由分发的key
	//	favo.Exchange,   // 绑定的exchange名
	//	false,         // 是否阻塞
	//	nil,           // 额外属性
	//)
	//if err != nil {
	//	fmt.Println("Bind Queue Failed", err)
	//	return
	//}
	// 发送消息
	errP := favo.Channel.Publish(
		favo.Exchange,   // 交换器名
		favo.RoutingKey, // routing key
		false,           // 是否返回消息(匹配队列)，如果为true, 会根据binding规则匹配queue，如未匹配queue，则把发送的消息返回给发送者
		false,           // 是否返回消息(匹配消费者)，如果为true, 消息发送到queue后发现没有绑定消费者，则把发送的消息返回给发送者
		amqp.Publishing{ // 消息内容
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	if errP != nil {
		log.Printf("Publish Message Failed", err)
		return
	}
}

// 消费者
func (favo *RabbitMQ) Consume() {
	// 声明队列
	_, err := favo.Channel.QueueDeclare(
		favo.QueueName,
		true,  // 是否持久化
		false, // 是否自动删除(前提是至少有一个消费者连接到这个队列，之后所有与这个队列连接的消费者都断开时，才会自动删除。注意：生产者客户端创建这个队列，或者没有消费者客户端与这个队列连接时，都不会自动删除这个队列)
		false, // 是否为排他队列（排他的队列仅对“首次”声明的conn可见[一个conn中的其他channel也能访问该队列]，conn结束后队列删除）
		false, // 是否阻塞
		nil,   //额外属性
	)
	if err != nil {
		log.Printf("Declare Queue Failed", err)
		return
	}
	// 接收消息
	messages, err := favo.Channel.Consume(
		favo.QueueName, // 队列名
		"",             // 消费者名，用来区分多个消费者，以实现公平分发或均等分发策略
		true,           // 是否自动应答
		false,          // 是否排他
		false,          // 是否接收只同一个连接中的消息，若为true，则只能接收别的conn中发送的消息
		true,           // 队列消费是否阻塞
		nil,            // 额外属性
	)
	if err != nil {
		log.Printf("Consume Message Failed", err)
		return
	}
	ch := make(chan int) //无缓冲区channel
	switch favo.QueueName {
	case "favoAdd": //点赞
		go favo.ConsumeFavoAdd(messages)
	case "favoDel": //取消赞
		go favo.ConsumeFavoDel(messages)
	}
	log.Printf("[*] Waiting for messagees,To exit press CTRL+C")
	<-ch //由协程从channel中pop一个值或阻塞
}

// 执行点赞操作的消费者
func (favo *RabbitMQ) ConsumeFavoAdd(messages <-chan amqp.Delivery) {
	for msg := range messages {
		//参数解析
		params := strings.Split(fmt.Sprintf("%s", msg.Body), " ")
		userId, _ := strconv.ParseInt(params[0], 10, 64)
		videoId, _ := strconv.ParseInt(params[1], 10, 64)
		//操作数据库 查询点赞记录并更新
		favoRecord, err := models.GetFavoRecord(uint64(userId), uint64(videoId))
		if err != nil {
			log.Printf("Get FavoRecord Failed")
			return
		} else {
			if favoRecord == (models.UserFavo{}) { // 数据库中不存在点赞记录
				err := models.CreateFavoRecord(uint64(userId), uint64(videoId))
				if err != nil {
					log.Printf("Create FavoRecord Failed")
					return
				}
			} else { // 数据库中存在点赞记录 更新状态为1
				err := models.UpdateFavoStatus(uint64(userId), uint64(videoId), 1)
				if err != nil {
					log.Printf("Update FavoRecord Failed")
					return
				}
			}
		}
	}
}

// 执行取消赞操作的消费者
func (favo *RabbitMQ) ConsumeFavoDel(messages <-chan amqp.Delivery) {
	for msg := range messages {
		//参数解析
		params := strings.Split(fmt.Sprintf("%s", msg.Body), " ")
		userId, _ := strconv.ParseInt(params[0], 10, 64)
		videoId, _ := strconv.ParseInt(params[1], 10, 64)
		//操作数据库 查询点赞记录并更新
		favoRecord, err := models.GetFavoRecord(uint64(userId), uint64(videoId))
		if err != nil {
			log.Printf("Get FavoRecord Failed")
			return
		} else {
			if favoRecord == (models.UserFavo{}) { // 数据库中不存在点赞记录
				log.Printf("Find FavoRecord Failed")
				return
			} else { // 数据库中存在点赞记录 更新状态为0
				err := models.UpdateFavoStatus(uint64(userId), uint64(videoId), 0)
				if err != nil {
					log.Printf("Update FavoRecord Failed")
					return
				}
			}
		}
	}
}
