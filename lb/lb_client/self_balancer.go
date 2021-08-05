package main

import "google.golang.org/grpc/balancer"

// SelfBalancer 自定义负载均衡算法
type SelfBalancer struct {
}

// 当连接状态发生改变的时候，会调用这个函数。如果这个错误返回的是ErrBadResolverState。
// ClientConn会调用名字解析中的ResolveNow函数,直到UpdateClientConnState返回一个nil为止，其他错误出现是不管的
func (sb *SelfBalancer) UpdateClientConnState(ccs balancer.ClientConnState) error {
	// ccs.ResolverState
	return nil
}

// UpdateSubConnState 当一个子连接改变的时候。会调用这个函数
func (sb *SelfBalancer) UpdateSubConnState(sc balancer.SubConn, scc balancer.SubConnState) {
	sc.Connect()

}

// ResolverError 当命名解析出错的时候，会调用这个函数
func (sb *SelfBalancer) ResolverError(err error) {

}

// Close 关闭该负载均衡器的时候，会调用该函数
func (sb *SelfBalancer) Close() {

}
