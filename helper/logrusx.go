package helper

import "github.com/sirupsen/logrus"

type TraceIdHook struct {
	TraceId string
}

func NewTraceIdHook(traceId string) logrus.Hook {
	return &TraceIdHook{TraceId: traceId}
}

func (h *TraceIdHook) Fire(entry *logrus.Entry) error {
	entry.Data["traceId"] = h.TraceId
	return nil
}

func (h *TraceIdHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
