package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"

	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
)

// NameResolver 命名解析
type NameResolverBuilder struct {
	schema string
	ctx    context.Context
	client *clientv3.Client
}

// NewResolverBuilder 新实例化一个名字解析器
func NewResolverBuilder(
	ctx context.Context,
	client *clientv3.Client,
	sm string,
) *NameResolverBuilder {
	return &NameResolverBuilder{schema: sm, ctx: ctx, client: client}
}

// Schema
func (nrb *NameResolverBuilder) Scheme() string {
	return nrb.schema
}

func (nrb *NameResolverBuilder) Build(
	target resolver.Target,
	cc resolver.ClientConn,
	opts resolver.BuildOptions,
) (resolver.Resolver, error) {
	res := &etcdResolver{
		target:  target,
		cc:      cc,
		etcdCli: nrb.client,
	}
	res.ctx, res.cancel = context.WithCancel(nrb.ctx)
	res.prefix = fmt.Sprintf("/%s/%s/", target.Scheme, target.Endpoint)
	res.start()
	return res, nil
}

// ETCDResolver 名字解析器
type etcdResolver struct {
	ctx     context.Context
	target  resolver.Target
	cc      resolver.ClientConn
	prefix  string
	etcdCli *clientv3.Client
	addrs   []resolver.Address
	cancel  context.CancelFunc
	sync.WaitGroup
}

func (r *etcdResolver) start() {
	log.Println("map[string]interface{}", r.target.Scheme, r.target.Endpoint)
	r._SchemeGet()
}

// Get 用于向ETCD中查询相关的服务列表
func (r *etcdResolver) _SchemeGet() {
	resp, err := r.etcdCli.Get(r.ctx, r.prefix, clientv3.WithPrefix())
	if err != nil {
		log.Fatalln(err)
	}
	for _, inst := range resp.Kvs {
		r._SchemeCreate(string(inst.Key))
	}
	if err := r.cc.UpdateState(resolver.State{Addresses: r.addrs}); err != nil {
		log.Println("error:", err.Error())
	}
}

//  key的样式为 /nrs/{schme}/{service_name}/ip:port
func (r *etcdResolver) _GetAddr(key string) string {
	return strings.TrimPrefix(key, r.prefix)
}

// _SchemeDel 删除
func (r *etcdResolver) _SchemeDel(key string) {
	addr := r._GetAddr(key)
	for index, inst := range r.addrs {
		if inst.Addr == addr {
			r.addrs = append(r.addrs[:index], r.addrs[index+1:]...)
			if err := r.cc.UpdateState(resolver.State{Addresses: r.addrs}); err != nil {
				log.Println("error:", err.Error())
			}
		}
	}
}

// _SchemeCreate 创建
func (r *etcdResolver) _SchemeCreate(key string) {
	addr := r._GetAddr(key)
	r.addrs = append(r.addrs, resolver.Address{
		Addr:       addr,
		ServerName: r.target.Endpoint,
	})
}

// _SchemeWatch 监听
func (r *etcdResolver) _SchemeWatch() {
	r.Add(1)
	defer r.Done()
	cWatchChan := r.etcdCli.Watch(
		r.ctx,
		r.prefix,
		clientv3.WithPrefix(),
	)
	for {
		select {
		case <-r.ctx.Done():
			return
		case events, ok := <-cWatchChan:
			if !ok {
				return
			}
			for _, ev := range events.Events {
				switch ev.Type {
				case clientv3.EventTypePut:
					if ev.IsCreate() {
						// 进入这里代表创建
						r._SchemeCreate(string(ev.Kv.Key))
					}
				case clientv3.EventTypeDelete:
					r._SchemeDel(string(ev.Kv.Key))
				}
			}
		}
	}
}

// ResolveNow 在GRPC调用不可用的时候，会调用该函数
func (*etcdResolver) ResolveNow(o resolver.ResolveNowOptions) {
	log.Println("resolveNow:", o)
}

func (r *etcdResolver) Close() {
	r.cancel()
	r.Wait()
	log.Println("调用释放函数")
}

func init() {
	// resolver.Register(NewResolverBuilder("grpc"))
}
