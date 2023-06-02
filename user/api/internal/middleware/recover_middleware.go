package middleware

import "net/http"

type RecoverMiddleware struct {
}

func NewRecoverMiddleware() *RecoverMiddleware {
	return &RecoverMiddleware{}
}

func (m *RecoverMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		next(w, r)
	}
}
