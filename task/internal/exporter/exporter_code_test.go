package exporter

import (
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/model/mongo"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/svc"
)

func TestExporterCode_ExcelAssembly(t *testing.T) {
	t.Run("TestName", func(t *testing.T) {
		var (
			lists = []mongo.ExportTask{{
				ID:           "1",
				ExportType:   1,
				ExportState:  1,
				RequestArgs:  `{"name":"hello"}`,
				TargetEmails: []string{"test@email.com"},
			}}
			ctx *svc.ServiceContext
		)
		for _, list := range lists {
			logrus.Infof("%+v", list)
			fact := newExportFact(ctx, &list) // 根据不同报表类型创建不同类型

			fact.LoadStatsRecords() // 获取统计数据
			if err := fact.Err(); err != nil {
				continue
			}
			fact.ExcelAssembly() // 组装Excel报表
			if err := fact.Err(); err != nil {
				continue
			}
			fact.UploadObs() // 上传到obs
			if err := fact.Err(); err != nil {
				continue
			}

			fact.PostEmail() // 发送邮件(发邮件错误可以忽略,)
			fact.AfterPost() // 更新数据库
		}
	})
}
