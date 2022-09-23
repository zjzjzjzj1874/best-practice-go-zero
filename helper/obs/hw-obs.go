// Package obs huawei object storage service
package obs

import (
	"fmt"
	"net/http"
	"time"

	"github.com/zjzjzjzj1874/huaweicloud-sdk-go-obs/obs"
)

// TODO 抽象一个对象,上传对象只需要一个reader,然后获取对象的访问url => 周末抽空实现

// ConfObs 华为对象存储配置
type ConfObs struct {
	AK       string // accessKey
	SK       string // secretKey
	Endpoint string
	Bucket   string `json:",default=bucket"`
	Path     string `json:",default=project"`
}

// 初始化自定义transport
var maxIdleConns = 1000
var maxConnsPerHost = 1000
var idleConnTimeout = 60
var transport = &http.Transport{
	MaxIdleConns:        maxIdleConns,
	MaxIdleConnsPerHost: maxIdleConns,
	MaxConnsPerHost:     maxConnsPerHost,
	IdleConnTimeout:     time.Second * time.Duration(idleConnTimeout),
}

// NewHWObsClient 创建ObsClient结构体
func NewHWObsClient(conf ConfObs) *obs.ObsClient {
	client, err := obs.New(conf.AK, conf.SK, conf.Endpoint, obs.WithHttpTransport(transport))
	if err != nil {
		panic(fmt.Sprintf("init obs client failure:%s", err.Error()))
	}

	return client
}
