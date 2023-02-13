package exporter

import (
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/constants/types"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/svc"

	"github.com/globalsign/mgo/bson"
	"github.com/sirupsen/logrus"
)

func Export(iCtx interface{}) {
	ctx, ok := iCtx.(*svc.ServiceContext)
	if !ok {
		logrus.Errorf("ctx type error:%T", iCtx)
		return
	}

	query := bson.M{"post_num": bson.M{"$lte": 5}, "export_state": bson.M{"$in": []types.ExportState{types.EXPORT_STATE__TODO, types.EXPORT_STATE__FAILURE}}} // 查询待导出和导出失败的任务
	lists, err := ctx.ExportTaskModel.List(query)
	if err != nil {
		logrus.Errorf("[报表导出任务]list ExportTaskModel failure:[query:%+v,err:%s]", query, err.Error())
		return
	}

	if len(lists) == 0 {
		return
	}

	logrus.Infof("[报表导出任务]导出:有%d条记录等待导出.", len(lists))
	for _, list := range lists {
		logrus.Infof("[报表导出任务]:%+v", list)
		fact := newExportFact(ctx, &list) // 根据不同报表类型创建不同类型

		fact.LoadStatsRecords() // 获取统计数据
		if err = fact.Err(); err != nil {
			continue
		}
		fact.ExcelAssembly() // 组装Excel报表
		if err = fact.Err(); err != nil {
			continue
		}
		fact.UploadObs() // 上传到obs
		if err = fact.Err(); err != nil {
			continue
		}

		fact.PostEmail() // 发送邮件(发邮件错误可以忽略,)
		fact.AfterPost() // 更新数据库
	}
}
