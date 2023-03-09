package contextx

import (
	"context"
	"time"
)

// WithValueOnly 返回仅保留value的新context, 不受父级context cancel控制
func WithValueOnly(ctx context.Context) context.Context {
	return ValueOnlyContext{ctx}
}

type ValueOnlyContext struct {
	context.Context
}

func (ValueOnlyContext) Done() <-chan struct{} {
	return nil
}

func (ValueOnlyContext) Deadline() (deadline time.Time, ok bool) {
	return
}

func (ValueOnlyContext) Err() error {
	return nil
}
