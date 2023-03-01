package tracex

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewGrpcExporter 定义一个exporter, 他使用grpc与宿主机上的collector进行通信
// Note: 注意此函数所依赖的grpc版本很高，与mongo客户端版本有可能有冲突
// endpoint: 值类似(otlp-agent.default:4317)
func NewGrpcExporter(ctx context.Context, endpoint string) *otlptrace.Exporter {
	conn, err := grpc.DialContext(ctx, endpoint,
		// 使用不加密的通信，可选择TLS
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		logrus.Fatalf("无法建立与grpc collector的链接: %v", err)
	}

	exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		logrus.Fatalf("New grpc collector failure: %v", err)
	}

	return exporter
}

// NewHttpExporter 定义一个exporter, 他使用http与宿主机上的collector进行通信
// endpoint:取值类似(otlp-agent.default:4318)
func NewHttpExporter(ctx context.Context, endpoint string) *otlptrace.Exporter {
	client := otlptracehttp.NewClient(
		otlptracehttp.WithEndpoint(endpoint),
		otlptracehttp.WithInsecure(),
	)

	exporter, err := otlptrace.New(ctx, client)
	if err != nil {
		logrus.Fatalf("New http collector failure: %v", err)
	}

	return exporter
}
