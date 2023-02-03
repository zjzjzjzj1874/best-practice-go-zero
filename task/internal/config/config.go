package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	Cron struct {
		TaskTestSpec string
	}
}
