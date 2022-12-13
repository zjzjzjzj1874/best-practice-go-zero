package email

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

// EmailConf 邮件配置配置
type EmailConf struct {
	Username string `json:",default="` // 服务用户名
	Password string `json:",default="` // 服务密码
	Host     string `json:",default="` // SMTP域名
	Port     int    `json:",default="` // 端口
	From     string `json:",default="` // 发件人
}

func (e EmailConf) PostEmail(msgUrl string, emails []string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "<"+e.From+">")                             // 别名
	m.SetHeader("To", emails...)                                    // 用户
	m.SetHeader("Subject", "收到一份邮件")                                // 主题
	m.SetBody("text/html", fmt.Sprintf("内容明细导出成功，下载地址：%s", msgUrl)) // 正文

	// 根据发送成功与失败来修改状态
	d := gomail.NewDialer(e.Host, e.Port, e.Username, e.Password)
	return d.DialAndSend(m)
}
