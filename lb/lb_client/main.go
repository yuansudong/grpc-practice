package main

import (
	"context"
	"log"
	"time"
	"unary/protocol/generate/gopb/bussiness/service/v1/pbs_demo"

	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

var scheme string = "grpc"
var b1domain string = "grpc:///b1.example.grpc.io"

var b2domain string = "grpc:///b2.example.grpc.io"

var etcdAddr []string = []string{"localhost:2379"}

var timeout time.Duration = 5 * time.Second

// DoRPCs 用于做
func DoRPCs(cc *grpc.ClientConn) {
	hwc := pbs_demo.NewDemoClient(cc)
	for i := 0; i < 10; i++ {
		DoParamGet(hwc)
		time.Sleep(time.Second)
	}
}

func DoParamGet(client pbs_demo.DemoClient) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	rsp, err := client.ParamGet(ctx, &pbs_demo.C2SBody{
		ParamOne: "1111",
	})
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("ParamGet:", rsp.Text)
}

func main() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints: etcdAddr,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
	resolver.Register(NewResolverBuilder(
		context.Background(),
		client,
		scheme,
	))
	b1Conn, err := grpc.Dial(
		b1domain,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	)
	if err != nil {
		log.Fatalln(err.Error())
	}
	DoRPCs(b1Conn)

	b1Conn1, err := grpc.Dial(
		b1domain,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	)
	if err != nil {
		log.Fatalln(err.Error())
	}
	DoRPCs(b1Conn1)
	log.Println("************************************************")

	b2Conn, err := grpc.Dial(
		b2domain,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalln(err.Error())
	}
	DoRPCs(b2Conn)
	time.Sleep(time.Minute)
}
