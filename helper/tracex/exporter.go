package tracex

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 定义一个exporter, 他使用grpc与宿主机上的collector进行通信
// 注意此函数所依赖的grpc版本很高，与mongo客户端版本有可能有冲突
func newGrpcExporter(ctx context.Context, svcName string) (*otlptrace.Exporter, error) {
	// 本地宿主机collector的grpc服务
	otlpAgentAddr := "otlp-agent.default:4317"
	conn, err := grpc.DialContext(ctx, otlpAgentAddr,
		// 使用不加密的通信，可选择TLS
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, fmt.Errorf("无法建立与grpc collector的链接: %w", err)
	}

	exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))

	return exporter, nil
}

// 定义一个exporter, 他使用http与宿主机上的collector进行通信
func newHttpExporter(ctx context.Context) (*otlptrace.Exporter, error) {
	// 本地宿主机collector的http服务
	otlpAgentAddr := "otlp-agent.default:4318"

	client := otlptracehttp.NewClient(
		otlptracehttp.WithEndpoint(otlpAgentAddr),
		otlptracehttp.WithInsecure(),
	)

	exporter, err := otlptrace.New(ctx, client)

	if err != nil {
		return nil, fmt.Errorf("无法建立与http collector的链接: %w", err)
	}

	return exporter, nil
}
