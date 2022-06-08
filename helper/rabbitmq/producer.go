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

func Init(ctx context.Context, conf Config) {
	Producer = make(chan PublishMetaData, 10)
	go newRabbitMQQueue(ctx, conf)
}

// NewRabbitMQQueue 设计:能够处理多个queueName
func newRabbitMQQueue(ctx context.Context, conf Config) {
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
		q, err := ch.QueueDeclare(
			conf.Name,
			false,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			fmt.Printf("QueueDeclare failure:[name:%s,err:%s]\n", conf.Name, err)
			return
		}

		go syncRabbitMQ(ctx, ch, &q, i)
	}

	select {
	case <-ctx.Done():
		fmt.Printf("recevice exit signal:bye-bye\n")
	}
}

func syncRabbitMQ(ctx context.Context, ch *amqp.Channel, queue *amqp.Queue, idx int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("recevice exit signal:bye-bye\n")
			return
		case meta := <-Producer:
			data, err := meta.Marshal()
			if err != nil {
				fmt.Printf("meta marshal failure:[meta:%+v,err:%s]\n", meta, err)
				continue
			}
			err = ch.Publish("", queue.Name, false, false, amqp.Publishing{
				ContentType: "application/json",
				Body:        data,
				Timestamp:   time.Now(),
			})
			if err != nil {
				fmt.Printf("Publish failure:[name:%s,data:%+v,err:%s]", queue.Name, meta, err.Error())
				continue
			}
			fmt.Printf("Publish success:[data:%+v,idx:%d]", meta, idx)
		}
	}
}
