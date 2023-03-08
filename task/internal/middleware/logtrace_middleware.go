package middleware

import "net/http"

type LogTraceMiddleware struct {
}

func NewLogTraceMiddleware() *LogTraceMiddleware {
	return &LogTraceMiddleware{}
}

func (m *LogTraceMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Passthrough to next handler if need
		next(w, r)
	}
}
