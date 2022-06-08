package rabbitmq

import (
	"context"
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

func init() {
	fmt.Println("init something in rabbitmq")
}

func InitConsumer(ctx context.Context, conf Config) {
	Producer = make(chan PublishMetaData, 10)
	go newConsumer(ctx, conf)
}

// newProducer 设计:能够起多个协程生产数据
func newConsumer(ctx context.Context, conf Config) {
	conn, err := amqp.Dial(conf.Addr)
	if err != nil {
		fmt.Printf("Dial failure:[addr:%s,err:%s]\n", conf.Addr, err.Error())
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("Get Channel failure:[err:%s]\n", err.Error())
		return
	}
	defer ch.Close()

	for i := 0; i < int(conf.Goroutine); i++ {
		deliveries, err := ch.Consume(
			conf.Name,
			"",
			true, // 自动确认,可以无需调用ch.Ack确认消息
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			fmt.Printf("Consume failure:[name:%s,err:%s]\n", conf.Name, err)
			return
		}
		go asyncConsumer(ctx, deliveries, i)
	}

	select {
	case <-ctx.Done():
		fmt.Printf("recevice exit signal:bye-bye\n")
	}
}

// asyncConsumer 消费者异步消费数据
func asyncConsumer(ctx context.Context, deliveries <-chan amqp.Delivery, idx int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("recevice exit signal:bye-bye\n")
			return
		case delivery := <-deliveries:
			fmt.Printf("consumerAt:%v,producerAt:%s", time.Now(), delivery.Timestamp)
			meta := PublishMetaData{}
			err := meta.Unmarshal(delivery.Body)
			if err != nil {
				fmt.Printf("meta Unmarshal failure:[meta:%+v,err:%s]\n", delivery.Body, err)
				continue
			}
			// todo do next logic,save in DB or do other thing
			fmt.Printf("Consumer success:[data:%+v,idx:%d]", meta, idx)
		}
	}
}
