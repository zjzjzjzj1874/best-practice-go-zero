package imp

import (
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/constants/errors"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/logic/imp"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/svc"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/types"
	"net/http"
)

func ImportHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImportReq
		if err := httpx.ParsePath(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		if err := httpx.ParseForm(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		logrus.Infof("req.id =================== %s", req.Id)
		if r.MultipartForm == nil || r.MultipartForm.File["file"] == nil || len(r.MultipartForm.File["file"]) == 0 {
			httpx.Error(w, errors.NewStatusError(errors.StatusBadRequestError).WithMsg("请传入文本校对自定义文本模板"))
			return
		}
		fileHeaders := r.MultipartForm.File["file"]

		file, err := fileHeaders[0].Open()
		if err != nil {
			httpx.Error(w, err)
			return
		}
		defer func() { _ = file.Close() }()

		excel, err := excelize.OpenReader(file)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		defer func() { _ = excel.Close() }()

		// 获取Sheet1上所有单元格 => 下面-1是减去表头
		rows, err := excel.GetRows("Sheet1")
		if err != nil {
			httpx.Error(w, err)
			return
		}
		for idx, row := range rows {
			if idx == 0 { // 表头忽略不计
				logrus.Warnf("表头忽略不计:%v", row)
				continue
			}
		}

		l := imp.NewImportLogic(r.Context(), svcCtx)
		resp, err := l.Import(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
