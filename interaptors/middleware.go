package main

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// StreamServiceLogger 流式调用logger.
func StreamServiceLogger() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {

		return nil
	}
}

// StreamClientLogger 客户单流式调用logger.
func StreamClientLogger() grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		cCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()
		md := metadata.Pairs("Authorization", "zxz.axs.dss", "k1", "v2", "k2", "v3")
		cCtx = metadata.NewOutgoingContext(cCtx, md)
		return streamer(cCtx, desc, cc, method, opts...)
	}
}

// UnaryServiceLogger 服务端UNARY模式下的logger。
func UnaryServiceLogger() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		return handler(ctx, req)
	}
}

func UnaryClientLogger() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
