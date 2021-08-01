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
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

type DemoImpl struct {
	*pbs_demo.UnimplementedDemoServer
}

func (di *DemoImpl) SsStream(cs *pbs_demo.C2SRouter, stream pbs_demo.Demo_SsStreamServer) error {

	defer log.Println("服务端单向流退出")
	sc := new(pbs_demo.S2CRouter)
	sc.Text = cs.Text
	for i := 0; i < 10; i++ {
		if err := stream.Send(sc); err != nil {
			log.Println("单向流:", err.Error())
			return err
		}
		time.Sleep(time.Second)
	}
	return nil

}

func (di *DemoImpl) CsStream(stream pbs_demo.Demo_CsStreamServer) error {
	sc := new(pbs_demo.S2CRouter)
	stream.Context()
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
	return stream.SendAndClose(sc)
}
func (di *DemoImpl) BIDStream(stream pbs_demo.Demo_BIDStreamServer) error {

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
				goto end
			}
		}
		time.Sleep(time.Second)
	}
end:
	return nil
}

func (di *DemoImpl) ParamGet(ctx context.Context, request *pbs_demo.C2SBody) (*pbs_demo.S2CBody, error) {
	return &pbs_demo.S2CBody{
		Text: "this greet from china",
	}, nil

}

func (di *DemoImpl) Hello(context.Context, *pbs_demo.C2SHello) (*pbs_demo.S2CHello, error) {

	return &pbs_demo.S2CHello{
		Reply: "了解",
	}, nil
}

var addr string
var service bool
var call string
var gateway_socket bool
var gateway_h2c bool

func main() {
	flag.StringVar(&addr, "addr", "172.81.209.185:5050", "--addr=172.81.209.185:5050")
	flag.BoolVar(&service, "service", false, "--service")
	flag.BoolVar(&gateway_socket, "gateway_socket", false, "--gateway_socket")
	flag.BoolVar(&gateway_h2c, "gateway_h2c", false, "--gateway_h2c")
	flag.StringVar(&call, "call", "Hello", "--call=Hello")
	flag.Parse()
	if service {
		RunService()
	} else if gateway_h2c {
		RunGatewayH2C()
	} else if gateway_socket {
		RunGatewaySocket()
	} else {
		RunClient()
	}
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
	case "Hello":
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		_, err := pbs_demo.NewDemoClient(conn).Hello(ctx, &pbs_demo.C2SHello{
			Text: "人生若如初见，何事秋风悲画扇！",
		})
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
	case "csstream":
		stream, err := pbs_demo.NewDemoClient(conn).CsStream(context.Background())
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
		log.Println(rsp.Text)
	case "ssstream":
		stream, err := pbs_demo.NewDemoClient(conn).SsStream(context.Background(), &pbs_demo.C2SRouter{Text: "客户端流模式"})
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
		defer stream.CloseSend()
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

	case "bidstream":
		stream, err := pbs_demo.NewDemoClient(conn).BIDStream(context.Background())
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
					return
				}
			}
			time.Sleep(time.Second)
		}

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

func GRPCHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	const Major = 2
	const ContentType = "Content-Type"
	const GRPC = "application/grpc"
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == Major && strings.Contains(
			r.Header.Get(ContentType),
			GRPC) {
			log.Println("this is h2")
			grpcServer.ServeHTTP(w, r)
		} else {
			log.Println("this is h1")
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

// RunGatewayH2C
func RunGatewayH2C() {
	h2mux := runtime.NewServeMux()
	demo_service := new(DemoImpl)
	if err := pbs_demo.RegisterDemoHandlerServer(
		context.Background(),
		h2mux,
		demo_service); err != nil {
		log.Fatalln(err.Error())
	}
	var opts []grpc.ServerOption = []grpc.ServerOption{
		grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
			log.Println("我是经过了的")
			return handler(ctx, req)
		}),
	}
	grpcServer := grpc.NewServer(opts...)
	pbs_demo.RegisterDemoServer(grpcServer, demo_service)
	http.ListenAndServe(":8081", GRPCHandlerFunc(grpcServer, h2mux))
}
