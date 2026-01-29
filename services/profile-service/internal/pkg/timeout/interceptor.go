package timeout

import (
	"context"
	"log/slog"
	"strings"

	"google.golang.org/grpc"
)

func unaryInterceptor(serviceName string, settingsProvider func() *ServiceClientSettings) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		settings := settingsProvider()

		ctx, cancel := newCtxWithTimeout(ctx, serviceName, method, settings)
		defer cancel()

		if _, ok := ctx.Deadline(); !ok {
			slog.Info("request without deadline occurred", "method", method)
		}

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func NewClientTimeoutInterceptor(target string) grpc.UnaryClientInterceptor {
	serviceName := serviceFromTarget(target)

	return unaryInterceptor(serviceName, getServiceClientTimeoutSettings)
}

func serviceFromTarget(target string) string {
	// dns:///analytic:50051 → analytic:50051
	if strings.Contains(target, ":///") {
		target = strings.Split(target, ":///")[1]
	}

	// analytic:50051 → analytic
	host := target
	if strings.Contains(target, ":") {
		host = strings.Split(target, ":")[0]
	}

	return host
}
