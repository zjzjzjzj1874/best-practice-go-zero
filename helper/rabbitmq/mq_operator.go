package rabbitmq

import (
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type Client interface {
	// 交换机模式
	// 发布
	ExchangeProduce(mqConfig *MQConfig, message interface{}) (err error)
	// 订阅
	ExchangeSubscribe(mqConfig *MQConfig, f Consumer) (err error)
}

type defaultClient struct {
	*MQConfig
}

func MustNewClient(url string) *defaultClient {
	confg := &MQConfig{url: url, publishRetryTimes: DEFAULT_PUBLISH_RETRYTIMES}
	return &defaultClient{MQConfig: confg}
}

// 默认的队列声明
func (c *defaultClient) myQueue(ch *amqp.Channel, queue string) (amqp.Queue, error) {
	return ch.QueueDeclare(
		queue,
		true,  /*是否持久化到硬盘*/
		false, /*没有消费者使用时自动删除*/
		false, /*排他性*/
		false, /*无需等待*/
		nil,   /*其他参数*/
	)
}

// 消费
func (c *defaultClient) Consumer(config *MQConfig, f Consumer) (err error) {

	defer func() {
		// if err := recover(); err != nil {
		logrus.Error("rabbitmq consume 断开,尝试重连....")
		time.Sleep(5 * time.Second)
		c.Consumer(config, f)
		// }
	}()

	conn, err := amqp.Dial(c.url)
	if err != nil {
		logrus.Errorf("rabbitmq dial fail,err:%v", err)
		return err
	}
	defer conn.Close()

	// 打开通道
	ch, err := conn.Channel()
	if err != nil {
		logrus.Errorf("rabbitmq open channel fail,err:%v", err)
		return err
	}
	defer ch.Close()

	// 创建队列
	q, err := c.myQueue(ch, config.Queue)
	if err != nil {
		return err
	}

	isExchangeMode := config.Exchange != ""
	// 使用交换机
	if isExchangeMode {
		changetype := amqp.ExchangeFanout // 默认
		if config.ChangeType != "" {
			changetype = config.ChangeType
		}
		// 创建交换机
		err = ch.ExchangeDeclare(
			config.Exchange, // 交换机
			changetype,      // 交换类型
			true,            // 是否持久化
			false,           // 自动删除
			false,           // 内部，不暴露给其他代理
			false,           // 无需等待
			nil,             // 参数
		)
		if err != nil {
			return err
		}

		// 绑定交换机
		err = ch.QueueBind(
			q.Name,
			config.Route,    // 目前路由键与队列一样
			config.Exchange, // 交换机
			false,           // 无需等待
			nil,             // 参数
		)
		if err != nil {
			return err
		}
	}

	message, err := ch.Consume(
		config.Queue,
		"",    /*消费者*/
		false, /*自动确认*/
		false, /*排他性*/
		false, /*非本地*/
		false, /*无需等待*/
		nil,   /*其他参数*/
	)
	if err != nil {
		logrus.Errorf("rabbitmq consume registry fail,err:%v", err)
		return err
	}

	// 消费者断线重连
	closeChan := make(chan *amqp.Error, 1)
	notifyClose := ch.NotifyClose(closeChan) // 消费者的channel产生错误可以捕捉
	for {
		select {
		case e := <-notifyClose:
			logrus.Errorf("chan通道错误,err:%v", e)
			close(closeChan)
			return
		case msg := <-message:
			go f(msg)
		}
	}

}

// 生产
func (c *defaultClient) Publisher(config *MQConfig, body interface{}) (err error) {

	var retry = func(config *MQConfig, body interface{}) (err error) {
		conn, err := amqp.Dial(c.url)
		if err != nil {
			logrus.Errorf("rabbitmq dial fail,err:%v", err)
			return err
		}
		defer conn.Close()

		// 打开通道
		ch, err := conn.Channel()
		if err != nil {
			logrus.Errorf("rabbitmq open channel fail,err:%v", err)
			return err
		}
		defer ch.Close()

		isExchangeMode := config.Exchange != ""

		// 使用交换机
		if isExchangeMode {
			changetype := amqp.ExchangeFanout // 默认
			if config.ChangeType != "" {
				changetype = config.ChangeType
			}
			// 创建交换机
			err = ch.ExchangeDeclare(
				config.Exchange, // 交换机
				changetype,      // 交换类型
				true,            // 是否持久化
				false,           // 自动删除
				false,           // 内部，不暴露给其他代理
				false,           // 无需等待
				nil,             // 参数
			)
			if err != nil {
				return err
			}
		}

		// 创建队列
		_, err = c.myQueue(ch, config.Queue)
		if err != nil {
			logrus.Errorf("rabbitmq queue declare fail,err:%v", err)
			return err
		}

		// 判断键,如果是交换机模式为路由键,否则为队列
		key := config.Queue
		if isExchangeMode {
			key = config.Route
		}

		// 可靠消息确认
		{
			logx.Infof("【%s】可靠消息确认。\n", key)
			if err := ch.Confirm(false); err != nil {
				logx.Errorf("【%s】channel无法支持可靠消息确认！！！。\n", key)
			} else {
				confirms := ch.NotifyPublish(make(chan amqp.Confirmation, 1))
				defer confirmOne(key, confirms)
			}
		}
		// 推送
		msgBody, _ := json.Marshal(body)
		return ch.Publish(
			config.Exchange,
			key,   // 路由键
			false, // true且未绑定与路由密钥匹配的队列时，发布可能无法传递
			false, // 同上
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        msgBody,
			},
		)
	}

	var retryErr error
	for i := 0; i < c.publishRetryTimes; i++ {
		retryErr = retry(config, body)
		if retryErr == nil {
			return nil
		}
		logrus.Errorf("rabbitmq publish fail,error:%v,retry %d....", retryErr, i+1)
	}

	return retryErr
}

// 队列消费者
type Consumer func(value amqp.Delivery)

// 队列模式=========================================
// 队列发送
func (c *defaultClient) QueuePush(queue, message string) (err error) {
	return c.Publisher(&MQConfig{Queue: queue}, message)
}

// 队列消费
func (c *defaultClient) QueueConsume(queue string, f Consumer) (err error) {
	return c.Consumer(&MQConfig{Queue: queue}, f)
}

// 发布订阅模式=========================================
// 发布订阅模式--发布
func (c *defaultClient) ExchangeProduce(mqConfig *MQConfig, message interface{}) (err error) {
	return c.Publisher(mqConfig, message)
}

// 发布订阅模式--订阅
func (c *defaultClient) ExchangeSubscribe(mqConfig *MQConfig, f Consumer) (err error) {
	return c.Consumer(mqConfig, f)
}
