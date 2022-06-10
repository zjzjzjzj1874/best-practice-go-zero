package helper

import (
	"fmt"
	"log"
	"net/http"
)

// PprofConf pprof参数设置
type PprofConf struct {
	Debug     bool `json:",optional"` // 调试模式是否开启
	DebugPort int  `json:",optional"` // 调试端口
}

// OpenPPROF 根据debug模式来选择是否开启pprof监测
func OpenPPROF(conf PprofConf) {
	if !conf.Debug {
		return
	}
	if conf.Debug && conf.DebugPort == 0 {
		fmt.Printf("can not open pprof due to port is 0.\n")
		return
	}

	go func() {
		fmt.Printf("listen on %d ...\n", conf.DebugPort)

		err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", conf.DebugPort), nil)
		if err != nil {
			log.Fatalf("open pprof failure:[err:%s]", err.Error())
		}
	}()
}
