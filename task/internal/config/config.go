package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper/email"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper/obs"
)

type Config struct {
	rest.RestConf

	Cron struct {
		TaskTestSpec   string `json:",default=0 0 1 * * *"`
		TaskExportSpec string `json:",default=*/30 * * * * *"` // 执行导出任务
	}

	CacheRedis cache.CacheConf
	//MysqlConf  mysql.MysqlConfig
	//MongoDB struct {
	//	URL string // MongoDB数据库链接url
	//}
	EmailConf email.EmailConf
	HwObs     obs.ConfObs

	Swagger     []byte `json:",optional"`
	SwaggerPath string `json:",default=/app/swagger.json"`
}
