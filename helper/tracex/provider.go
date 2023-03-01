package tracex

import (
	"context"
	"sync"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

// GlobalTracer 相当于一个工厂函数
// 使用他将生产一个全局的tracer，可以把这个tracer当成logger一样的使用
var (
	once         sync.Once
	GlobalTracer trace.Tracer
)

func NewGrpcProvider(ctx context.Context, conf Trace) trace.Tracer {
	if !conf.isOpen() {
		return nil
	}

	once.Do(func() {
		exporter := NewGrpcExporter(ctx, conf.Endpoint)
		res, err := resource.New(
			context.Background(),
			resource.WithAttributes(semconv.ServiceNameKey.String(conf.Name)),
		)
		if err != nil {
			logrus.Fatalf("NewTraceProvider:%v", err)
		}

		// Register the trace exporter with a TracerProvider, using a batch
		// span processor to aggregate spans before export.
		bsp := sdktrace.NewBatchSpanProcessor(exporter)
		tracerProvider := sdktrace.NewTracerProvider(
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithResource(res),
			sdktrace.WithSpanProcessor(bsp),
		)
		// 全局设置
		otel.SetTracerProvider(tracerProvider)

		// 这个设置可以开启服务间的上下文传递.
		otel.SetTextMapPropagator(propagation.TraceContext{})

		// Initialize global tracer
		GlobalTracer = tracerProvider.Tracer("GlobalTracer")
	})

	return GlobalTracer
}

func NewHttpProvider(ctx context.Context, conf Trace) trace.Tracer {
	if !conf.isOpen() {
		return nil
	}

	once.Do(func() {
		exporter := NewHttpExporter(ctx, conf.Endpoint)
		res, err := resource.New(
			context.Background(),
			resource.WithAttributes(semconv.ServiceNameKey.String(conf.Name)),
		)
		if err != nil {
			logrus.Fatalf("NewTraceProvider:%v", err)
		}

		// Register the trace exporter with a TracerProvider, using a batch
		// span processor to aggregate spans before export.
		bsp := sdktrace.NewBatchSpanProcessor(exporter)
		tracerProvider := sdktrace.NewTracerProvider(
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithResource(res),
			sdktrace.WithSpanProcessor(bsp),
		)
		// 全局设置
		otel.SetTracerProvider(tracerProvider)

		// 这个设置可以开启服务间的上下文传递.
		otel.SetTextMapPropagator(propagation.TraceContext{})

		// Initialize global tracer
		GlobalTracer = tracerProvider.Tracer("GlobalTracer")
	})

	return GlobalTracer
}
