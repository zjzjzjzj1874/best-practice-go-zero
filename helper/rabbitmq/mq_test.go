package rabbitmq

import (
	"fmt"
	"log"
	"testing"

	"github.com/streadway/amqp"
)

var client *defaultClient

func init() {
	client = MustNewClient("amqp://guest:guest@127.0.0.1:5672/").WithPublishRetryTimes(3)
}

// 队列模式
func Test_Queue(t *testing.T) {
	// 发布
	if err := client.QueuePush("test-queue", "测试消息"); err != nil {
		log.Println(err)
	}

	// 消费
	_ = client.QueueConsume("test-queue", func(value amqp.Delivery) {
		fmt.Println("我是来消费的", value)
	})
}

// 发布订阅模式
func Test_Subscribe(t *testing.T) {
	mqConfig := &MQConfig{
		Queue:      "test-queue",
		Route:      "test-route",
		Exchange:   "test-Exchange",
		ChangeType: "topic",
	}
	// 订阅1
	_ = client.ExchangeSubscribe(mqConfig, func(value amqp.Delivery) {
		fmt.Println("我是订阅者1，接收消息：", value)
	})

	// 订阅2
	_ = client.ExchangeSubscribe(mqConfig, func(value amqp.Delivery) {
		fmt.Println("我是订阅者2，接收消息：", value)
	})

	// 发送1
	_ = client.ExchangeProduce(mqConfig, "这是发布者发送的消息")
}

// 路由模式
func Test_Route(t *testing.T) {
	// 队列1,订阅route1消息
	_ = client.Consumer(&MQConfig{Exchange: "direct-exchange", Queue: "direct-queue1", Route: "route1", ChangeType: "direct"}, func(value amqp.Delivery) {
		fmt.Println("我是队列1,接收route1的消息:", value.Body)
	})

	// 队列2,订阅route1消息
	_ = client.Consumer(&MQConfig{Exchange: "direct-exchange", Queue: "direct-queue2", Route: "route1", ChangeType: "direct"}, func(value amqp.Delivery) {
		fmt.Println("我是队列2,接收route1的消息:", value.Body)
	})

	// 队列3,订阅route2消息
	_ = client.Consumer(&MQConfig{Exchange: "direct-exchange", Queue: "direct-queue3", Route: "route2", ChangeType: "direct"}, func(value amqp.Delivery) {
		fmt.Println("我是队列3,接收route2的消息:", value.Body)
	})

	// 发送到route1
	_ = client.Publisher(&MQConfig{Exchange: "direct-exchange", Route: "route1", ChangeType: "direct"}, "这是发送给route1的消息")
}
