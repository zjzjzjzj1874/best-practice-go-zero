package helper

type Swagger struct {
	Path string `json:",default=/app/swagger.json"` // swagger.json的默认位置
	Data []byte `json:",optional"`                  // swagger的数据,可以通过go:embed嵌入后赋值
}
