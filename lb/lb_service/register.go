package main

import (
	"fmt"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"golang.org/x/net/context"
)

type EtcdRegister struct {
	_prefix string
}

// New 用于实例化。
func New(scheme, domain, addr string) *EtcdRegister {
	return &EtcdRegister{
		_prefix: fmt.Sprintf("/%s/%s/%s", scheme, domain, addr),
	}
}

func (er *EtcdRegister) Do(ctx context.Context, client *clientv3.Client) error {
	sess, err := concurrency.NewSession(client, concurrency.WithTTL(1200))
	if err != nil {
		return err
	}
	_, err = client.Put(ctx, er._prefix, "", clientv3.WithLease(sess.Lease()))
	if err != nil {
		return err
	}
	return err
}
