package main

import (
	"fmt"
	"helloworld/internal/conf"

	"github.com/hashicorp/consul/api"

	"github.com/go-kratos/kratos/contrib/config/consul/v2"
	"github.com/go-kratos/kratos/v2/config"
)

func main() {
	consulClient, err := api.NewClient(&api.Config{
		Address: "192.168.10.210:8500",
	})
	if err != nil {
		panic(err)
	}
	//consul中需要标注文件后缀，kratos读取配置需要适配文件后缀
	//eg:key设置为Dev_Auth.yaml
	cs, err := consul.New(consulClient, consul.WithPath("Dev_Auth"))

	// The file suffix needs to be marked, and kratos needs to adapt the file suffix to read the configuration.
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
	fmt.Println(bc.Server)
}
