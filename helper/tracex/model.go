package tracex

// Trace 配置
type Trace struct {
	Name     string `json:",optional"`
	Endpoint string `json:",optional"`
	Batcher  string `json:",default=jaeger,options=jaeger|zipkin|otel"`
}

// 是否开启链路追踪
func (c Trace) isOpen() bool {
	return len(c.Endpoint) != 0
}
