package tracex

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
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

// base = 基础类key
const (
	// DBSystemKey is the attribute Key conforming to the "db.system" semantic conventions
	DBSystemKey = attribute.Key("db.system")
	// MiddlewareTypeKey is the attribute Key conforming to the "middleware.type" semantic conventions
	MiddlewareTypeKey = attribute.Key("middleware.type")
	// DevopsTypeKey is the attribute Key conforming to the "devops.type" semantic conventions
	DevopsTypeKey = attribute.Key("devops.type")
)

// business = 业务类Key
const (
	// TaskKey is the attribute Key conforming to the "task" semantic conventions,especially for service task
	TaskKey = attribute.Key("task")
	// MyZeroKey is the attribute Key conforming to the "myzero" semantic conventions,especially for service my-zero
	MyZeroKey = attribute.Key("myzero")
)

var (
	// DBSystemPostgreSQL postgresql
	DBSystemPostgreSQL = DBSystemKey.String("postgresql")
	// DBSystemMySQL MySQL
	DBSystemMySQL = DBSystemKey.String("mysql")

	// MiddlewareElasticSearch es
	MiddlewareElasticSearch = MiddlewareTypeKey.String("elasticsearch")
	// MiddlewareKafka kafka
	MiddlewareKafka = MiddlewareTypeKey.String("kafka")
	// MiddlewareRabbitmq rabbitmq
	MiddlewareRabbitmq = MiddlewareTypeKey.String("rabbitmq")
	// MiddlewareRedis redis
	MiddlewareRedis = MiddlewareTypeKey.String("redis")

	// DevopsElasticSearch es
	DevopsElasticSearch = DevopsTypeKey.String("elasticsearch")
	// DevopsKibana kibana
	DevopsKibana = DevopsTypeKey.String("kibana")
	// DevopsFilebeat filebeat
	DevopsFilebeat = DevopsTypeKey.String("filebeat")
	// DevopsPrometheus prometheus
	DevopsPrometheus = DevopsTypeKey.String("prometheus")
	// DevopsGrafana grafana
	DevopsGrafana = DevopsTypeKey.String("grafana")
	// DevopsJaeger jaeger
	DevopsJaeger = DevopsTypeKey.String("jaeger")
	// DevopsOpentelemetry opentelemetry
	DevopsOpentelemetry = DevopsTypeKey.String("opentelemetry")
)

// Task returns an attribute KeyValue conforming to the
// "task" semantic conventions. It represents the
// identifies the context in which an task happened.
func Task(val string) attribute.KeyValue {
	return TaskKey.String(val)
}

// MyZero returns an attribute KeyValue conforming to the
// "myzero" semantic conventions. It represents the
// identifies the context in which a zero happened.
func MyZero(val string) attribute.KeyValue {
	return MyZeroKey.String(val)
}
