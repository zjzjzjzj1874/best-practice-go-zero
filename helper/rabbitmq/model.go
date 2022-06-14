package rabbitmq

import jsoniter "github.com/json-iterator/go"

type Config struct {
	// 消费者
	Consumer struct {
		Addr string // amqp地址
		Name string // 队列名称
	}
	// 生产者
	Producer struct {
		Addr string // amqp地址
		Name string // 队列名称
	}
}

// // 全局生产者:为什么需要私有化全局生产者:防止produce被到处乱传,造成意外,所以需要到处一个共有方法来处理
var produce chan PublishMetaData

func ProduceData(meta PublishMetaData) {
	produce <- meta
}

type PublishMetaData struct {
	Name    string   `json:"name"`    // 姓名
	Age     int      `json:"age"`     // 年龄
	Hobbies []string `json:"hobbies"` // 爱好
}

func (p *PublishMetaData) Marshal() ([]byte, error) {
	return jsoniter.Marshal(p)
}

func (p *PublishMetaData) Unmarshal(data []byte) error {
	return jsoniter.Unmarshal(data, p)
}
