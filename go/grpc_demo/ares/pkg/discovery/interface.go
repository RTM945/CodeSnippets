package discovery

import "context"

// ServiceInstance 服务描述
type ServiceInstance struct {
	ID      string
	Name    string
	Address string
}

// ServiceRegistry 注册/注销
type ServiceRegistry interface {
	Register(ctx context.Context, inst ServiceInstance) error
	Deregister(ctx context.Context, inst ServiceInstance) error
}

// ServiceResolver List/Watch
type ServiceResolver interface {
	List(ctx context.Context, serviceName string) ([]ServiceInstance, error)
	Watch(ctx context.Context, serviceName string) (<-chan []ServiceInstance, error)
}
