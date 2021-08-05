package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
	"unary/protocol/generate/gopb/bussiness/service/v1/pbs_demo"

	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
)

type DemoImpl struct {
	*pbs_demo.UnimplementedDemoServer
	addr string
}

// NewDemo 新建立一个Demo
func NewDemo(ar string) *DemoImpl {
	return &DemoImpl{addr: ar}
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

// ParamGet
func (di *DemoImpl) ParamGet(ctx context.Context, request *pbs_demo.C2SBody) (*pbs_demo.S2CBody, error) {
	return &pbs_demo.S2CBody{
		Text: "this greet from china	" + di.addr,
	}, nil

}

//  Hello Hello 接口
func (di *DemoImpl) Hello(context.Context, *pbs_demo.C2SHello) (*pbs_demo.S2CHello, error) {

	return &pbs_demo.S2CHello{
		Reply: "了解	" + di.addr,
	}, nil
}

var scheme string
var domain string
var addr string

func main() {
	flag.StringVar(&scheme, "scheme", "grpc", "--scheme=grpc")
	flag.StringVar(&addr, "addr", "localhost:5051", "--addr=localhost:5051")
	flag.StringVar(&domain, "domain", "b1.example.grpc.io", "--addr=b1.example.grpc.io")
	flag.Parse()
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}

	New(scheme, domain, addr).Do(context.Background(), client)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err.Error())
		return
	}
	s := grpc.NewServer()
	pbs_demo.RegisterDemoServer(s, NewDemo(addr))
	fmt.Printf("serving on %s\n", addr)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
		return
	}
}
