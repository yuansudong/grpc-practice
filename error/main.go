package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
	"unary/protocol/generate/gopb/bussiness/service/v1/pbs_demo"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type DemoImpl struct {
	*pbs_demo.UnimplementedDemoServer
}

func (di *DemoImpl) ParamGet(ctx context.Context, request *pbs_demo.C2SBody) (*pbs_demo.S2CBody, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		for _, v := range md {
			log.Println(strings.Join(v, ","))
		}
	}
	nmd := metadata.New(map[string]string{
		"x-forward-for": "192.168.2.222",
	})
	mmd := metadata.Pairs("k1", "v1", "k2", "v2")
	nmd.Set("level-key", "level-val")
	nmd.Append("level-key", "level-val2")
	log.Println(nmd.Copy().Get("x-forward-for"))
	grpc.SetHeader(ctx, metadata.Join(md, nmd, mmd))
	grpc.SetTrailer(ctx, metadata.New(map[string]string{
		"x-trial-forward-for": "192.168.2.222",
		"x-trial-position":    "这里代表trial的位置",
	}))
	log.Println(metadata.Join(md, nmd, mmd).Len())
	return &pbs_demo.S2CBody{
		Text: "this greet from china",
	}, status.New(codes.NotFound, "this is test error").Err()

}

func (di *DemoImpl) CsStream(stream pbs_demo.Demo_CsStreamServer) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if ok {
		for _, v := range md {
			log.Println(strings.Join(v, ","))
		}
	}
	sc := new(pbs_demo.S2CRouter)
	sc.Text = "客户端单向流"
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		log.Println(req.Text)
	}
	nmd := metadata.New(map[string]string{
		"x-cs-forward-for": "192.168.2.222",
	})
	mmd := metadata.Pairs("csk1", "v1", "k2", "v2")
	nmd.Set("level-cs-key", "level-val")
	nmd.Append("level-cs-key", "level-val2")
	stream.SetHeader(metadata.Join(md, nmd, mmd))
	stream.SetTrailer(metadata.Join(md, nmd, mmd))
	return stream.SendAndClose(sc)
}

func (di *DemoImpl) SsStream(cs *pbs_demo.C2SRouter, stream pbs_demo.Demo_SsStreamServer) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if ok {
		for _, v := range md {
			log.Println(strings.Join(v, ","))
		}
	}
	sc := new(pbs_demo.S2CRouter)
	sc.Text = cs.Text
	for i := 0; i < 10; i++ {
		if err := stream.Send(sc); err != nil {
			log.Println("单向流:", err.Error())
			return err
		}
		time.Sleep(time.Second)
	}

	nmd := metadata.New(map[string]string{
		"x-ss-forward-for": "192.168.2.222",
	})
	mmd := metadata.Pairs("ssk1", "v1", "k2", "v2")
	nmd.Set("level-ss-key", "level-val")
	nmd.Append("level-ss-key", "level-val2")
	stream.SetHeader(metadata.Join(md, nmd, mmd))
	stream.SetTrailer(metadata.Join(md, nmd, mmd))
	return nil
}

func (di *DemoImpl) BIDStream(stream pbs_demo.Demo_BIDStreamServer) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if ok {
		log.Println(md)
	}
	go func() {
		rsp, err := stream.Recv()
		if err != nil {
			return
		}
		log.Println(rsp.Text)
		time.Sleep(time.Second)
	}()
	for {
		err := stream.Send(&pbs_demo.S2CRouter{Text: "客户端流模式"})
		if err != nil {
			if err != nil {
				return err
			}
		}
		time.Sleep(time.Second)
	}
}

func (di *DemoImpl) Hello(context.Context, *pbs_demo.C2SHello) (*pbs_demo.S2CHello, error) {
	return &pbs_demo.S2CHello{
		Reply: "了解",
	}, nil
}

var addr string
var service bool
var call string
var gateway bool

func main() {
	flag.StringVar(&addr, "addr", "172.81.209.185:5050", "--addr=172.81.209.185:5050")
	flag.BoolVar(&service, "service", false, "--service")
	flag.BoolVar(&gateway, "gateway", false, "--gateway")
	flag.StringVar(&call, "call", "Hello", "--call=Hello")
	flag.Parse()
	if service {
		RunService()
	} else if gateway {
		RunGatewaySocket()
	} else {
		RunClient()
	}

}

// CallBIDStream
func CallBIDStream(conn *grpc.ClientConn) {
	md := metadata.Pairs("k1", "v1", "k1", "v2", "k2", "v3")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	stream, err := pbs_demo.NewDemoClient(conn).BIDStream(ctx)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	defer stream.CloseSend()
	go func() {
		rsp, err := stream.Recv()
		if err != nil {
			return
		}
		log.Println(rsp.Text)
		time.Sleep(time.Second)
	}()
	for {
		err := stream.Send(&pbs_demo.C2SRouter{Text: "客户端流模式"})
		if err != nil {
			if err != nil {
				goto end
			}
		}
		time.Sleep(time.Second)
	}
end:
	stream.CloseSend()
	log.Println(stream.Header())
	log.Println(stream.Trailer())
}

// CallSsStream 调用流模式
func CallSsStream(conn *grpc.ClientConn) {
	md := metadata.Pairs("k1", "v1", "k1", "v2", "k2", "v3")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	stream, err := pbs_demo.NewDemoClient(conn).SsStream(ctx, &pbs_demo.C2SRouter{Text: "客户端流模式"})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	defer stream.CloseSend()
	log.Println(stream.Header())
	for i := 0; i < 10; i++ {
		rsp, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err.Error())
		}
		log.Println(rsp.Text)
	}
	log.Println(stream.Trailer())
}

// CallCsStream 调用客户端流模式
func CallCsStream(conn *grpc.ClientConn) {
	md := metadata.Pairs("k1", "v1", "k1", "v2", "k2", "v3")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	stream, err := pbs_demo.NewDemoClient(conn).CsStream(ctx)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	for i := 0; i < 10; i++ {
		stream.Send(&pbs_demo.C2SRouter{Text: "客户端流模式"})
	}
	rsp, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	log.Println(stream.Header())
	log.Println(stream.Trailer())
	log.Println(rsp.Text)
}

func CallParamGet(conn *grpc.ClientConn) {
	// 创建一个新的context
	ctx := metadata.AppendToOutgoingContext(context.Background(), "hello-key", "hello-val")
	recvHeader := metadata.New(map[string]string{})
	recvTrial := metadata.New(map[string]string{})
	rsp, err := pbs_demo.NewDemoClient(conn).ParamGet(
		ctx,
		&pbs_demo.C2SBody{},
		grpc.Header(&recvHeader),
		grpc.Trailer(&recvTrial),
	)
	if sts := status.FromContextError(err); sts != nil {
		log.Println(sts.Code(), sts.Message())
		os.Exit(-1)
	}
	fmt.Println(rsp.Text, recvHeader, recvTrial)
}

func RunClient() {
	var opts []grpc.DialOption = []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	defer conn.Close()
	switch call {
	case "ParamGet":
		CallParamGet(conn)
	case "CsStream":
		CallCsStream(conn)
	case "SsStream":
		CallSsStream(conn)
	case "BidStream":
		CallBIDStream(conn)
	}
}

// RunService
func RunService() {
	lis, err := net.Listen("tcp", ":5050")
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		os.Exit(-1)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pbs_demo.RegisterDemoServer(grpcServer, new(DemoImpl))
	grpcServer.Serve(lis)
}

// gateway的方式
func RunGatewaySocket() {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pbs_demo.RegisterDemoHandlerFromEndpoint(context.Background(), mux, "localhost:5050", opts)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	http.ListenAndServe(":8081", mux)
}
