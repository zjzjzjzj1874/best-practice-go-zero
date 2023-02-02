package middlewares

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper"
)

type LogTraceMiddleware struct {
}

func NewLogTraceMiddleware() *LogTraceMiddleware {
	return &LogTraceMiddleware{}
}

const TraceId = "traceId"

func (m *LogTraceMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		traceID := r.Header.Get(TraceId)
		if traceID == "" {
			uid := uuid.NewString()
			r.Header.Set(TraceId, traceID)
			traceID = uid
		}
		logrus.AddHook(helper.NewTraceIdHook(traceID))
		w.Header().Set(TraceId, traceID)
		next(w, r)
	}
}
