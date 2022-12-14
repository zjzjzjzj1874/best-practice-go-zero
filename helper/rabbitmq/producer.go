package rabbitmq

import (
	"context"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

func init() {
	//fmt.Println("init something in rabbitmq Producer")
}

func InitProducer(ctx context.Context, conf Config) {
	produce = make(chan PublishMetaData, 10)
	go newProducer(conf)
}

// newProducer 设计:能够起多个协程生产数据
func newProducer(conf Config) {
	conn, err := amqp.Dial(conf.Producer.Addr)
	if err != nil {
		fmt.Printf("Dial failure:[conf:%+v,err:%s]\n", conf.Producer, err.Error())
		log.Fatal(err)
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

	q, err := ch.QueueDeclare(
		conf.Producer.Name,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Printf("QueueDeclare failure:[name:%s,err:%s]\n", conf.Producer.Name, err)
		return
	}

	cond := true
	for cond {
		select {
		case e := <-notifyClose:
			fmt.Printf("receive producer chan err:[err:%s]\n", e.Error())
			//close(notifyClose)

			newProducer(conf) // 断线之后重连
			cond = false
		case meta := <-produce:
			data, err := meta.Marshal()
			if err != nil {
				fmt.Printf("meta marshal failure:[meta:%+v,err:%s]\n", meta, err)
				continue
			}
			err = ch.Publish("", q.Name, false, false, amqp.Publishing{
				ContentType: "application/json",
				Body:        data,
				Timestamp:   time.Now(),
			})
			if err != nil {
				fmt.Printf("Publish failure:[name:%s,data:%+v,err:%s]\n", q.Name, meta, err.Error())
				continue
			}
			fmt.Printf("Publish success:[data:%+v]\n", meta)
		}
	}
}
