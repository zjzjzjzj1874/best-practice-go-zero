package helper

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

func InitLogrus() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:   "2006-01-02T15:04:05.000Z",
		DisableTimestamp:  false,
		DisableHTMLEscape: true,
		DataKey:           "",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "msg",
			logrus.FieldKeyFunc:  "caller",
		},
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			return "", strings.Replace(frame.Function, filepath.Ext(frame.Function),
				fmt.Sprintf("/%s:%d", filepath.Base(frame.File), frame.Line), -1)
		},
	})
}

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
