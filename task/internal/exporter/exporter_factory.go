package exporter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/constants/types"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper/obs"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/model/mongo"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/svc"
)

// ExporterFact 导出接口类型
type ExporterFact interface {
	LoadStatsRecords() // 0.获取统计数据
	ExcelAssembly()    // 1.组装Excel报表
	UploadObs()        // 2.上传obs成功后设置链接
	PostEmail()        // 3.获取到链接后发邮件
	AfterPost()        // 4.钩子函数,完成后触发
	Err() error        // 获取err
}

type Exporter struct {
	err      error         // 是否有错误
	emailUrl string        // 邮件url
	buf      *bytes.Buffer // 邮件内容
	ctx      *svc.ServiceContext
	task     *mongo.ExportTask
	records  *[]interface{} // 等待处理的数据
}

func NewExport(ctx *svc.ServiceContext, task *mongo.ExportTask) *Exporter {
	return &Exporter{
		ctx:  ctx,
		task: task,
	}
}

func (e *Exporter) Err() error {
	return e.err
}

// PostEmail 发送邮件
func (e *Exporter) PostEmail() {
	if e.emailUrl == "" || len(e.task.TargetEmails) == 0 {
		e.task.Msg = fmt.Sprintf("无需发邮件:emailUrl:%s,emails:%v", e.emailUrl, e.task.TargetEmails)
		e.task.ExportState = types.EXPORT_STATE__SUCCESS
		return
	}
	e.task.PostNum++ // 发送邮件次数++

	if err := e.ctx.Config.EmailConf.PostEmail(e.emailUrl, e.task.TargetEmails); err != nil {
		e.err = err
		logrus.Errorf("PostEmail_Fialure:[url:%s,emails:%v,err:%s]", e.emailUrl, e.task.TargetEmails, err.Error())
		e.task.Msg = "发送邮件失败"
		e.task.ExportState = types.EXPORT_STATE__FAILURE
		return
	}

	e.task.Msg = "邮件发送成功"
	e.task.ExportState = types.EXPORT_STATE__SUCCESS
}

// UploadObs 上传obs
func (e *Exporter) UploadObs() {
	if e.buf == nil || e.buf.Len() == 0 {
		return
	}
	uid := uuid.NewString()
	e.emailUrl, e.err = e.ctx.HwObsClient.PutObjectAndSetMetadata(e.buf, fmt.Sprintf("%s-%s.%s", e.task.ExportType.Label(), uid, REPORT_FORMAT_XLSX), ContentTypes[REPORT_FORMAT_XLSX], &obs.SetObjectMetaData{
		ContentDisposition: "attachment",
	})
	if e.err != nil {
		logrus.Errorf("PutObjectAndSetMetadata failure:[err:%s]", e.err.Error())
	} else {
		e.task.EmailUrl = e.emailUrl
	}
}

// AfterPost 发送后处理逻辑:更新数据库
func (e *Exporter) AfterPost() {
	e.task.UpdateTime = time.Now()
	e.err = e.ctx.ExportTaskModel.Update(e.task)
	if e.err != nil {
		logrus.Errorf("Update_ExportTask_Failure:[task:%+v,err:%s]", e.task, e.err)
	}
}

// LoadStatsRecords 获取统计数据
func (e *Exporter) LoadStatsRecords() {
	req := ExportReq{} // 获取查询数据
	e.err = json.Unmarshal([]byte(e.task.RequestArgs), &req)
	if e.err != nil {
		logrus.Errorf("json.Unmarshal Failure:[reqArgs:%s,err:%s]", e.task.RequestArgs, e.err)
		return
	}
	// 参数校验
	startTime, err := time.Parse(yyyyMMDD, req.StartDate)
	if err != nil {
		e.err = err
		logrus.Errorf("start_date参数格式错误,example'2006-01-02':%s", req.StartDate)
		return
	}
	endTime, err := time.Parse(yyyyMMDD, req.EndDate)
	if err != nil {
		e.err = err
		logrus.Errorf("end_date参数格式错误,example'2006-01-02':%s", req.EndDate)
		return
	}

	// 生成筛选条件
	query := bson.M{"id": req.Id}
	statStartTime := time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 0, 0, 0, 0, time.UTC)
	statEndTime := time.Date(endTime.Year(), endTime.Month(), endTime.Day(), 23, 59, 59, 59, time.UTC)
	query["stat_time"] = bson.M{"$gte": statStartTime, "$lte": statEndTime}

	// TODO find your own data
}

// 工厂类方法获取
func newExportFact(ctx *svc.ServiceContext, task *mongo.ExportTask) ExporterFact {
	var exporter ExporterFact
	switch task.ExportType {
	case types.EXPORT_TYPE__CODE:
		exporter = NewExporterCode(ctx, task)
	case types.EXPORT_TYPE__REPORT:
		exporter = NewExporterReport(ctx, task)
	}

	return exporter
}
