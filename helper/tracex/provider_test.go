package tracex

import (
	"context"
	"errors"
	"testing"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func Test_NewHttpProvider(t *testing.T) {
	type args struct {
		ctx  context.Context
		conf Trace
	}
	tests := []struct {
		name string
		args args
		want trace.Tracer
	}{
		{name: "Test_HTTP_Trace", args: args{
			ctx: context.Background(),
			conf: Trace{
				Name:     "TestTracingCode",
				Endpoint: "localhost:4318",
				Batcher:  "",
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tracer := NewHttpProvider(tt.args.ctx, tt.args.conf)
			ctx, span := tracer.Start(tt.args.ctx, "TestNewHttpProvider")
			// 可以给该span附加键值对信息
			span.SetAttributes(attribute.String("NewHttpProvider", "Test"))
			defer span.End()

			callChild(ctx)
			time.Sleep(time.Second * 5)
		})
	}
}

func callChild(ctx context.Context) {

	_, span := GlobalTracer.Start(ctx, "TestNewHttpProviderChild")
	// 可以给该span附加键值对信息
	span.SetAttributes(attribute.String("NewHttpProviderChild", "TestChild"))
	defer span.End()

	span.RecordError(errors.New("这是一个测试的错误"))
}
