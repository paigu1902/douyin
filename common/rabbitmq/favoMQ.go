package rabbitmq

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/streadway/amqp"
	"paigu1902/douyin/common/models"
	"strconv"
	"strings"
)

var RmqFavoAdd *RabbitMQ
var RmqFavoDel *RabbitMQ

// 初始化RabbitMQ
func init() {
	RmqFavoAdd = InitRabbitMQ("favoAdd")
	go RmqFavoAdd.Consume()
	RmqFavoDel = InitRabbitMQ("favoDel")
	go RmqFavoDel.Consume()
}

// Publish 生产者
func (favo *RabbitMQ) Publish(msg string) {
	// 1. 声明队列
	q, err := favo.Channel.QueueDeclare(
		favo.QueueName,
		true,  // 是否持久化
		false, // 是否自动删除(前提是至少有一个消费者连接到这个队列，之后所有与这个队列连接的消费者都断开时，才会自动删除。注意：生产者客户端创建这个队列，或者没有消费者客户端与这个队列连接时，都不会自动删除这个队列)
		false, // 是否为排他队列（排他的队列仅对“首次”声明的conn可见[一个conn中的其他channel也能访问该队列]，conn结束后队列删除）
		false, // 是否阻塞
		nil,   //额外属性
	)
	if err != nil {
		panic(err)
		return
	}
	// 2. 发送消息
	errP := favo.Channel.Publish(
		"",     // 交换器名
		q.Name, // routing key
		false,  // 是否返回消息(匹配队列)，如果为true, 会根据binding规则匹配queue，如未匹配queue，则把发送的消息返回给发送者
		false,  // 是否返回消息(匹配消费者)，如果为true, 消息发送到queue后发现没有绑定消费者，则把发送的消息返回给发送者
		amqp.Publishing{ // 消息内容
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	if errP != nil {
		klog.Info("Publish Message Failed", err)
		return
	}
	return
}

// Consume 消费者
func (favo *RabbitMQ) Consume() {
	// 1. 声明队列
	_, err := favo.Channel.QueueDeclare(
		favo.QueueName,
		true,  // 是否持久化
		false, // 是否自动删除(前提是至少有一个消费者连接到这个队列，之后所有与这个队列连接的消费者都断开时，才会自动删除。注意：生产者客户端创建这个队列，或者没有消费者客户端与这个队列连接时，都不会自动删除这个队列)
		false, // 是否为排他队列（排他的队列仅对“首次”声明的conn可见[一个conn中的其他channel也能访问该队列]，conn结束后队列删除）
		false, // 是否阻塞
		nil,   //额外属性
	)
	if err != nil {
		klog.Info("Declare Queue Failed", err)
		return
	}
	// 2. 接收消息
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
		klog.Info("Consume Message Failed", err)
		return
	}
	klog.Info("Consume Consume")
	ch := make(chan int) //无缓冲区channel
	switch favo.QueueName {
	case "favoAdd": //点赞
		go favo.ConsumeFavoAdd(messages)
	case "favoDel": //取消赞
		go favo.ConsumeFavoDel(messages)
	default:
		klog.Info("RabbitMQ Actiontype Error")
	}
	klog.Info("[*] Waiting for messagees,To exit press CTRL+C")
	<-ch //由协程从channel中pop一个值或阻塞
}

// ConsumeFavoAdd 执行点赞操作的消费者
func (favo *RabbitMQ) ConsumeFavoAdd(messages <-chan amqp.Delivery) {
	for msg := range messages {
		// 1. 参数解析
		params := strings.Split(fmt.Sprintf("%s s b add ", msg.Body), " ")
		userId, _ := strconv.ParseInt(params[0], 10, 64)
		videoId, _ := strconv.ParseInt(params[1], 10, 64)
		// 2. 操作数据库 查询点赞记录并更新
		favoRecord, err1 := models.GetFavoRecord(uint64(userId), uint64(videoId))
		if err1 != nil && err1.Error() != "record not found" {
			klog.Info("ConsumeFavoAdd Get FavoRecord Failed")
			continue
		}
		// 3. 若数据库中不存在点赞记录 创建记录
		if favoRecord == (models.UserFavo{}) {
			record := models.UserFavo{UserId: uint64(userId), VideoId: uint64(videoId), Status: 1}
			err2 := models.CreateFavoRecord(&record)
			if err2 != nil {
				klog.Info("Create FavoRecord Failed")
			}
		} else { // 4. 若数据库中存在点赞记录 更新状态为1
			req := models.UserFavo{UserId: uint64(userId), VideoId: uint64(videoId), Status: 1}
			err3 := models.UpdateFavoStatus(&req)
			if err3 != nil {
				klog.Info("Update FavoRecord Failed")
			}
		}
	}
}

// ConsumeFavoDel 执行取消赞操作的消费者
func (favo *RabbitMQ) ConsumeFavoDel(messages <-chan amqp.Delivery) {
	for msg := range messages {
		// 1. 参数解析
		params := strings.Split(fmt.Sprintf("%s", msg.Body), " ")
		userId, _ := strconv.ParseInt(params[0], 10, 64)
		videoId, _ := strconv.ParseInt(params[1], 10, 64)
		// 2. 操作数据库 查询点赞记录并更新
		favoRecord, err1 := models.GetFavoRecord(uint64(userId), uint64(videoId))
		if err1 != nil && err1.Error() != "record not found" {
			klog.Info("Get FavoRecord Failed")
			continue
		}
		// 3. 若数据库中不存在点赞记录
		if favoRecord == (models.UserFavo{}) {
			klog.Info("Find FavoRecord Failed")
		} else { // 4. 若数据库中存在点赞记录 更新状态为0
			req := models.UserFavo{UserId: uint64(userId), VideoId: uint64(videoId), Status: 0}
			err2 := models.UpdateFavoStatus(&req)
			if err2 != nil {
				klog.Info("Update FavoRecord Failed")
			}
		}
	}
}
