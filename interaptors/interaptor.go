package main

import (
	"context"

	"google.golang.org/grpc"
)

// ChainUnaryClient 客户端拦截器调用链
func ChainUnaryClient(interceptors ...grpc.UnaryClientInterceptor) grpc.UnaryClientInterceptor {
	n := len(interceptors)

	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		chainer := func(currentInter grpc.UnaryClientInterceptor, currentInvoker grpc.UnaryInvoker) grpc.UnaryInvoker {
			return func(currentCtx context.Context, currentMethod string, currentReq, currentRepl interface{}, currentConn *grpc.ClientConn, currentOpts ...grpc.CallOption) error {
				return currentInter(currentCtx, currentMethod, currentReq, currentRepl, currentConn, currentInvoker, currentOpts...)
			}
		}
		chainedInvoker := invoker
		for i := n - 1; i >= 0; i-- {
			chainedInvoker = chainer(interceptors[i], chainedInvoker)
		}

		return chainedInvoker(ctx, method, req, reply, cc, opts...)
	}
}

// ChainUnaryServer UNARY模式下服务端拦截器。
func ChainUnaryServer(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	n := len(interceptors)
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		chainer := func(currentInter grpc.UnaryServerInterceptor, currentHandler grpc.UnaryHandler) grpc.UnaryHandler {
			return func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
				return currentInter(currentCtx, currentReq, info, currentHandler)
			}
		}
		chainedHandler := handler
		for i := n - 1; i >= 0; i-- {
			chainedHandler = chainer(interceptors[i], chainedHandler)
		}
		return chainedHandler(ctx, req)
	}
}

// ChainStreamClient 客户端链式调用
func ChainStreamClient(interceptors ...grpc.StreamClientInterceptor) grpc.StreamClientInterceptor {
	n := len(interceptors)
	if n == 0 {
		return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
			return streamer(ctx, desc, cc, method, opts...)
		}
	}
	if n == 1 {
		return interceptors[0]
	}
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		currStreamer := streamer
		for i := n - 1; i > 0; i-- {
			innerStreamer, i := currStreamer, i
			currStreamer = func(currentCtx context.Context, currentDesc *grpc.StreamDesc, currentConn *grpc.ClientConn, currentMethod string, currentOpts ...grpc.CallOption) (grpc.ClientStream, error) {
				return interceptors[i](currentCtx, currentDesc, currentConn, currentMethod, innerStreamer, currentOpts...)
			}
		}
		return interceptors[0](ctx, desc, cc, method, currStreamer, opts...)
	}
}

// ChainStreamServer  服务端的流式调用链
func ChainStreamServer(interceptors ...grpc.StreamServerInterceptor) grpc.StreamServerInterceptor {
	n := len(interceptors)
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		chainer := func(currentInter grpc.StreamServerInterceptor, currentHandler grpc.StreamHandler) grpc.StreamHandler {
			return func(currentSrv interface{}, currentStream grpc.ServerStream) error {
				return currentInter(currentSrv, currentStream, info, currentHandler)
			}
		}
		chainedHandler := handler
		for i := n - 1; i >= 0; i-- {
			chainedHandler = chainer(interceptors[i], chainedHandler)
		}
		return chainedHandler(srv, ss)
	}
}
