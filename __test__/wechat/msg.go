package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/eatmoreapple/openwechat"
	"github.com/sirupsen/logrus"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper"
)

var f = flag.String("f", "", "传入文件地址")
var t = flag.Int64("t", 0, "请传入执行时间")
var g = flag.String("g", "", "请传入群组名称")

// 使用方法: go run . -f ~/a.txt -t 1692698540 -g "测试群"  => 在1692698540这个时间点,读取a.txt的文件内容,然后向测试群发送消息.  运行程序之前要先登录微信桌面版哦

// TODO 待发送完成后,把发送的数据写到一个归档的文件里面,然后清空这一个
func main() {
	flag.Parse()
	if *f == "" || *t == 0 || *g == "" || *t < time.Now().Unix() {
		logrus.Errorf("文件:%s,时间:%d,群组:%s", *f, *t, *g)
		return
	}

	initLogrus()
	logrus.Infof("发送群组:%s,文件:%s,执行时间:%d,%v", *g, *f, *t, time.Unix(*t, 0).Local().Format(time.RFC3339))

	con, err := ioutil.ReadFile(*f)
	if err != nil {
		logrus.Errorf("读取文件失败,[txt:%s,err:%s]", *f, err.Error())
		return
	}
	if len(con) == 0 {
		logrus.Errorf("空文件,[txt:%s]", *f)
		return
	}

	logrus.Infof("读取消息,[con:%s]", con)

	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式

	// 注册消息处理函数
	bot.MessageHandler = func(msg *openwechat.Message) {
		if msg.IsText() && msg.Content == "ping" {
			msg.ReplyText("pong")
		}
	}
	// 注册登陆二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	// 登陆
	if err := bot.Login(); err != nil {
		logrus.Println(err)
		return
	}

	// 获取登陆的用户
	self, err := bot.GetCurrentUser()
	if err != nil {
		logrus.Println(err)
		return
	}

	// 获取所有的好友
	//friends, err := self.Friends()
	//fmt.Println(friends, err)

	// 获取所有的群组
	groups, err := self.Groups()
	if err != nil {
		logrus.Errorf("Groups,[err:%s]", err.Error())
		return
	}

	ng := groups.SearchByNickName(1, *g)
	for i := 0; i < ng.Count(); i++ {
		g := ng[i]
		logrus.Infof("Group == %+v", g)
	}
	logrus.Infof("找出群聊:%+v", ng)

	done := make(chan bool) // 退出信号

	// 计算距离指定时间点的时间间隔
	duration := time.Until(time.Unix(0, 0))
	time.AfterFunc(duration, func() {
		err = ng.SendText(string(con), time.Second)
		if err != nil {
			logrus.Errorf("发送消息失败,[err:%s]", err.Error())
			return
		}

		logrus.Infof("发送成功")
		done <- true

	})

	//go func(g2 openwechat.Groups) {
	//	for {
	//		logrus.Infof("读取消息,[con:%s]", con)
	//		time.Sleep(time.Second * 10)
	//		err = g2.SendText(string(con), time.Second)
	//		if err != nil {
	//			logrus.Errorf("发送消息失败,[err:%s]", err.Error())
	//			return
	//		}
	//	}
	//}(ng)

	<-done
	logrus.Infof("退出")
	bot.Exit()

	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	//bot.Block()
}

func initLogrus() {
	helper.InitLogrus()

	log, err := os.OpenFile(fmt.Sprintf("%s.log", *f), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		logrus.Errorf("open %s file failed [error: %s]", *f, err.Error())
		return
	}

	logrus.SetOutput(log)
}
