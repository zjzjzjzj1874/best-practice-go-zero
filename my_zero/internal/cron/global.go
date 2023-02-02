package cron

import (
	"github.com/sirupsen/logrus"
	"github.com/zeromicro/go-zero/core/threading"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/svc"
)

var globalMap = make(map[string]interface{}) // 全局map,用来触发调用

// 将每个任务注册到全局map中
func set(name string, value interface{}) {
	val, ok := globalMap[name]
	if ok && val != value {
		logrus.Warnf("global Map is diff from cache:old:%v,new:%v", val, value)
		return
	}
	globalMap[name] = value
}

func RunWithName(svcCtx *svc.ServiceContext, name string) string {
	val, ok := globalMap[name]
	if !ok {
		logrus.Errorf("不存在改任务:%s", name)
		return "请先注册任务"
	}
	switch tt := val.(type) {
	case *task:
		threading.GoSafe(func() {
			tt.do(svcCtx)
		})
	default:
		return "未知类型"
	}

	return "执行成功"
}
