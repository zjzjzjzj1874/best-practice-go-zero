package middlewares

import (
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper/tracex"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/semconv/v1.18.0/httpconv"
	"go.opentelemetry.io/otel/trace"
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

		// 初始化链路追踪,从HTTP HEADER中取出上游TRACE信息，加入上下文
		newCtx := otel.GetTextMapPropagator().Extract(
			r.Context(), propagation.HeaderCarrier(r.Header),
		)
		r.WithContext(newCtx)
		var span trace.Span
		newCtx, span = tracex.GlobalTracer.Start(newCtx, "Log Trace Handle", trace.WithNewRoot(),
			trace.WithAttributes(httpconv.ServerRequest(r.URL.Path, r)...))
		defer span.End()

		logrus.AddHook(helper.NewTraceIdHook(traceID))
		w.Header().Set(TraceId, traceID)
		next(w, r)
	}
}
