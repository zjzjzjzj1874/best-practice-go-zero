package rabbitmq

const DEFAULT_PUBLISH_RETRYTIMES = 3 // 默认推送重试次数

// 内部调用配置，也可以供外部实现需求
// Exchange：交换机名 QueueRoute：队列/路由键名 ChangeType：交换类型
// 1、exchange不为空
// 		1、如果路由键为空实现发布订阅模式（具体还取决于changeType）
//		2、如果路由键不为空则为路由模式（具体还取决于changeType）
// 2、exchange为空队列模式
type MQConfig struct {
	Queue, Route      string
	Exchange          string // 交换机名
	ChangeType        string // 交换类型 faout、direct、header
	publishRetryTimes int    // 推送失败重试次数
	url               string // 连接url
}

// 设置推送重试次数
func (client *defaultClient) WithPublishRetryTimes(times int) *defaultClient {
	client.MQConfig.publishRetryTimes = times
	return client
}
