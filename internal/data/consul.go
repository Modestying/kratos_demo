package data

import (
	"helloworld/internal/conf"
	"time"

	kratos_consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/hashicorp/consul/api"
)

func NewConsulConfig(conf *conf.Data_Consul, s *conf.Service, server *conf.Server) (*kratos_consul.Registry, func(), error) {
	consulConfig := &api.Config{Address: conf.Addr, WaitTime: time.Second * 10}
	client, err := api.NewClient(consulConfig)
	if err != nil {
		panic("Generate consul client")
	}
	opts := []kratos_consul.Option{
		kratos_consul.WithHeartbeat(false),
		kratos_consul.WithHealthCheck(true),
		kratos_consul.WithHealthCheckInterval(30),
		kratos_consul.WithDeregisterCriticalServiceAfter(5),
	}
	consulClient := kratos_consul.New(client, opts...)

	//此处进行自定义服务注册
	//如server/WebServer,属于自定义服务，无法自动注册到consul中，需要手动注册/注销
	//svc := &registry.ServiceInstance{
	//	Name:      "web",          //服务名
	//	ID:        "web" + s.Node, //服务ID
	//	Version:   "0.1",          //版本
	//	Endpoints: []string{"http://" + server.Web.Addr},
	//}

	////注册web服务
	//consulClient.Register(context.Background(), svc)
	cleanup := func() {
		//consulClient.Deregister(context.Background(), svc)
	}
	return consulClient, cleanup, nil
}
