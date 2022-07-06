package main

import (
	"fmt"
	config_consul "github.com/go-kratos/kratos/contrib/config/consul/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/hashicorp/consul/api"
	"helloworld/internal/conf"
)

func main() {

	consulClient, err := api.NewClient(&api.Config{
		Address: "192.168.10.153:8500",
	})
	if err != nil {
		panic(err)
	}
	cs, err := config_consul.New(consulClient, config_consul.WithPath("app/Auth.yaml"))
	//consul中需要标注文件后缀，kratos读取配置需要适配文件后缀
	//The file suffix needs to be marked, and kratos needs to adapt the file suffix to read the configuration.
	if err != nil {
		panic(err)
	}
	c := config.New(config.WithSource(cs))
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	fmt.Println(bc.Server.Http.Addr)
}
