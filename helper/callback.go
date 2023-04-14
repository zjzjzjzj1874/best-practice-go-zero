package helper

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/constants/errors"
	"net/http"
)

func UnAuthCallback(w http.ResponseWriter, r *http.Request, err error) {
	if err != nil {
		err = errors.NewUnauthorizedErr(err.Error())
		httpx.Error(w, err)
	}
}
