package tracex

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

// TraceProvider相当于一个工厂函数
// 使用他将生产一个全局的tracer，可以把这个tracer当成logger一样的使用
var GlobalTracer trace.Tracer

func newTraceProvider(exp *otlptrace.Exporter) {

	res, err := resource.New(context.Background(),
		resource.WithAttributes(
			// 该服务名会展示在UI上
			semconv.ServiceNameKey.String("服务名"),
		),
	)

	if err != nil {
		panic(err)
	}
	// Register the trace exporter with a TracerProvider, using a batch
	// span processor to aggregate spans before export.
	bsp := sdktrace.NewBatchSpanProcessor(exp)
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

}
