package tracex

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

// InjectToMsg traceCarrier可以放进队列消息体中
func InjectToMsg(ctx context.Context) propagation.MapCarrier {
	traceCarrier := propagation.MapCarrier{} // map[string]string
	otel.GetTextMapPropagator().Inject(ctx, traceCarrier)

	return traceCarrier
}

// ExtractFromMsg 抽取trace信息, 放入ctx
func ExtractFromMsg(msg propagation.MapCarrier) context.Context {
	ctx := otel.GetTextMapPropagator().Extract(
		context.Background(), msg,
	)
	return ctx
}

// ExtractFromMsgWithCtx 抽取trace信息, 放入ctx
func ExtractFromMsgWithCtx(ctx context.Context, msg propagation.MapCarrier) context.Context {
	ctx = otel.GetTextMapPropagator().Extract(
		ctx, msg,
	)
	return ctx
}
