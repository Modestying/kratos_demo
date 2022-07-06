package main

import (
	"context"
	registry "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/hashicorp/consul/api"
	"log"
	"time"
)

import (
	. "helloworld/api/helloworld/v1"
)

func main() {
	conf := &api.Config{Address: "192.168.10.153:8500"}
	client, err := api.NewClient(conf)
	if err != nil {
		panic(err)
	}

	dis := registry.New(client)
	//<schema>://[namespace]/<service-name>
	//namespace是consul 企业版才有 https://www.consul.io/commands/namespace#basic-examples
	//default 是默认值.最下面：https://www.consul.io/api-docs/discovery-chain
	endpoint := "discovery://default/Demo"
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint(endpoint), grpc.WithDiscovery(dis))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	dd := NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	name := "this is WIN"
	r, err := dd.SayHello(ctx, &HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
