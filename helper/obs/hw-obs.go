// Package obs huawei object storage service
package obs

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/zjzjzjzj1874/huaweicloud-sdk-go-obs/obs"

	"github.com/sirupsen/logrus"
	"github.com/zeromicro/go-zero/rest"
)

// ConfObs 华为对象存储配置
type ConfObs struct {
	AK          string // accessKey
	SK          string // secretKey
	Endpoint    string
	Bucket      string `json:",default=bucket"`
	Path        string `json:",default=project"`
	ServiceName string `json:",optional"` // 子服务名称
	Mode        string `json:",optional"`
	Project     string `json:",default="`     // 项目
	URLTimeout  int    `json:",default=3600"` // URL的超时时间，默认1个小时
}

type HwObsClient struct {
	*obs.ObsClient
	ConfObs
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
func NewHWObsClient(conf ConfObs, restConf rest.RestConf) *HwObsClient {
	if restConf.Mode == "dev" { // 开发环境不注册
		return &HwObsClient{}
	}
	client, err := obs.New(conf.AK, conf.SK, conf.Endpoint, obs.WithHttpTransport(transport))
	if err != nil {
		panic(fmt.Sprintf("init obs client failure:%s", err.Error()))
	}

	return (&HwObsClient{
		ObsClient: client,
		ConfObs:   conf,
	}).WithServiceIfAbsent(restConf.Name).WithModeIfAbsent(restConf.Mode)
}

// WithModeIfAbsent mode初始化
func (c *HwObsClient) WithModeIfAbsent(mode string) *HwObsClient {
	if c.Mode == "" {
		c.Mode = mode
	}
	return c
}

// WithServiceIfAbsent serviceName初始化
func (c *HwObsClient) WithServiceIfAbsent(service string) *HwObsClient {
	if c.ServiceName == "" {
		c.ServiceName = service
	}
	return c
}

// GetObsSignUrl 获取带认证url
func (c *HwObsClient) GetObsSignUrl(key string) string {
	url, err := c.getObsSignUrl(key)
	if err != nil {
		return ""
	}
	return url
}

func (c *HwObsClient) getObsSignUrl(key string) (string, error) {
	input := &obs.CreateSignedUrlInput{}
	input.Bucket = c.Bucket
	input.Key = key
	input.Method = obs.HttpMethodGet
	input.Expires = c.URLTimeout // 图片URL超时时间 => 暂定1小时
	output, err := c.CreateSignedUrl(input)
	if err != nil {
		if obsError, ok := err.(obs.ObsError); ok {
			logrus.Errorf("CreateSignedUrl failure :[key:%s,Code:%s,Message:%s]", key, obsError.Code, obsError.Message)
			return "", obsError
		}
		logrus.Errorf("CreateSignedUrl failure :[object:%s,err:%s]", input.Key, err.Error())
		return "", err
	}

	return output.SignedUrl, nil
}

// SetObjectMetaData 修改对象header中的metadata
type SetObjectMetaData struct {
	CacheControl       string
	ContentDisposition string
	ContentEncoding    string
	ContentLanguage    string
	ContentType        string
	Expires            string
	Metadata           map[string]string
}

// PutObjectAndSetMetadata 上传并修改header相关header元数据
func (c *HwObsClient) PutObjectAndSetMetadata(file io.Reader, objectID, contentType string, meta *SetObjectMetaData) (string, error) {
	obsInput := &obs.PutObjectInput{}
	obsInput.Bucket = c.Bucket
	// 注意:这里key的规则 => 项目/环境(dev|test|pro)/服务/对象名(ObjectID)
	obsInput.Key = fmt.Sprintf("%s/%s/%s/%s", c.Project, c.Mode, c.ServiceName, objectID)
	obsInput.Body = file
	obsInput.ContentType = contentType
	obsInput.ContentDisposition = meta.ContentDisposition

	_, err := c.ObsClient.PutObject(obsInput)
	if err != nil {
		logrus.Errorf("PutObject failure[key:%s,%s]", obsInput.Key, err.Error())
		return "", err
	}

	return c.getObsSignUrl(obsInput.Key)
}
