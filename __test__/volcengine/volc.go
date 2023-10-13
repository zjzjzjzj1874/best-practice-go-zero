/*
 * File: volc.go
 * Created Date: 2023-10-13 10:14:02
 * Author: zjzjzjzj1874
 * Description:  火山引擎文本翻译 https://www.volcengine.com/docs/4640/65067
 */
package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	kAccessKey      = "密钥管理中的 AccessKey ID" // https://console.volcengine.com/iam/keymanage/
	kSecretKey      = "密钥管理中的 AccessKey Secret"
	kServiceVersion = "2020-06-01"
)

var (
	ServiceInfo = &base.ServiceInfo{
		Timeout: 5 * time.Second,
		Host:    "open.volcengineapi.com",
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
		Credentials: base.Credentials{Region: base.RegionCnNorth1, Service: "translate"},
	}
	ApiInfoList = map[string]*base.ApiInfo{
		"TranslateText": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"TranslateText"},
				"Version": []string{kServiceVersion},
			},
		},
	}
)

func newClient() *base.Client {
	client := base.NewClient(ServiceInfo, ApiInfoList)
	client.SetAccessKey(kAccessKey)
	client.SetSecretKey(kSecretKey)
	return client
}

// ----------
type transReq struct {
	TargetLanguage string   `json:"TargetLanguage"`
	TextList       []string `json:"TextList"`
}
type TranslationList struct {
	Translation            string `json:"Translation"`
	DetectedSourceLanguage string `json:"DetectedSourceLanguage"`
}

type transResp struct {
	TranslationList  []TranslationList     `json:"TranslationList"`
	ResponseMetaData base.ResponseMetadata `json:"ResponseMetaData"`
}

type Translator struct {
	client *base.Client
}

func NewTranslator() *Translator {
	return &Translator{
		client: newClient(),
	}
}

// 中译英
func (t *Translator) TransZhToEn(zh string) string {
	if zh == "" {
		return ""
	}
	req := transReq{
		TargetLanguage: "en",
		TextList:       []string{zh},
	}
	body, err := json.Marshal(req)
	if err != nil {
		logx.Error(err)
		return zh
	}
	resp, _, err := t.client.Json("TranslateText", nil, string(body))
	if err != nil {
		logx.Error(err)
		return zh
	}

	r := new(transResp)
	err = json.Unmarshal(resp, &r)
	if err != nil {
		logx.Error(err)
		return zh
	}

	if len(r.TranslationList) > 0 {
		return r.TranslationList[0].Translation
	}
	logx.Infof("【volc 翻译】%#v", r.ResponseMetaData.Error)

	return ""
}
