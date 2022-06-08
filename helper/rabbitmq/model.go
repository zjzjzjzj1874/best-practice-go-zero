package rabbitmq

import jsoniter "github.com/json-iterator/go"

type Config struct {
	Addr      string
	Name      string // 队列名称
	Goroutine uint   // 协程数量
}

var Producer chan PublishMetaData // 生产者

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
