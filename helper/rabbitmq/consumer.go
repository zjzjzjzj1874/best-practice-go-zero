package rabbitmq

import (
	"context"
	"fmt"
	"github.com/streadway/amqp"
	"time"
)

func init() {
	fmt.Println("init something in rabbitmq consumer")
}

func InitConsumer(ctx context.Context, conf Config) {
	go newConsumer(conf)
}

// newProducer 设计:能够起多个协程生产数据
func newConsumer(conf Config) {
	conn, err := amqp.Dial(conf.Consumer.Addr)
	if err != nil {
		fmt.Printf("Dial failure:[consumer:%+v,err:%s]\n", conf.Consumer, err.Error())
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("Get Channel failure:[err:%s]\n", err.Error())
		return
	}
	defer ch.Close()

	closeChan := make(chan *amqp.Error, 1)
	notifyClose := ch.NotifyClose(closeChan)

	deliveries, err := ch.Consume(
		conf.Consumer.Name,
		"",
		true, // 自动确认,可以无需调用ch.Ack确认消息
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Printf("Consume failure:[name:%s,err:%s]\n", conf.Consumer.Name, err)
		return
	}

	cond := true
	for cond {
		select {
		case e := <-notifyClose:
			fmt.Printf("receive chan err:[err:%s]\n", e.Error())
			close(notifyClose)

			newConsumer(conf) // 断线之后重连
			cond = false
		case delivery := <-deliveries:
			fmt.Printf("consumerAt:%v,producerAt:%s\n", time.Now(), delivery.Timestamp)
			meta := PublishMetaData{}
			err := meta.Unmarshal(delivery.Body)
			if err != nil {
				fmt.Printf("meta Unmarshal failure:[meta:%+v,err:%s]\n", delivery.Body, err)
				continue
			}
			// todo do next logic,save in DB or do other thing
			fmt.Printf("Consumer success:[data:%+v]\n", meta)
		}
	}
}
