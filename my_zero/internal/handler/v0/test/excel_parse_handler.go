package test

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/logic/v0/test"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/svc"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/types"
	"net/http"
	"strings"
)

func ExcelParseHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.String()
		if strings.LastIndex(path, "?") != -1 {
			path = path[:strings.LastIndex(path, "?")]
			path = path[strings.Index(path, "/")+1:]
		}
		fmt.Println("path === ", path)

		var req types.ListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		if err := helper.Validate().StructCtx(r.Context(), &req); err != nil {
			httpx.Error(w, err)
			return
		}

		fileHeaders := r.MultipartForm.File["file"]
		if fileHeaders == nil || len(fileHeaders) == 0 {
			httpx.Error(w, errors.New("请传入文件"))
			return
		}

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
		// 获取 Sheet1 上所有单元格
		rows, err := excel.GetRows("Sheet1")
		for idx, row := range rows {
			if idx == 0 {
				continue // skip table header
			}
			for _, colCell := range row {
				fmt.Print(colCell, "\t")
			}
		}

		l := test.NewExcelParseLogic(r.Context(), svcCtx)
		resp, err := l.ExcelParse(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
