package main

import (
	"context"

	"log"
	"time"

	. "helloworld/api/helloworld/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// conf := &api.Config{Address: "192.168.10.153:8500"}
	// client, err := api.NewClient(conf)
	// if err != nil {
	// 	panic(err)
	// }

	// dis := registry.New(client)
	// //<schema>://[namespace]/<service-name>
	// //namespace是consul 企业版才有 https://www.consul.io/commands/namespace#basic-examples
	// //default 是默认值.最下面：https://www.consul.io/api-docs/discovery-chain
	// endpoint := "discovery://default/Demo"
	// conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint(endpoint), grpc.WithDiscovery(dis))
	// if err != nil {
	// 	panic(err)
	// }
	// defer conn.Close()

	// dnsServer := "0.0.0.0"
	// dnsPort := "8600"
	// serverHost := "hello.service.consul"
	// serverPort := "9000"
	// dialer := &net.Dialer{
	// 	LocalAddr: &net.UDPAddr{
	// 		IP: net.IPv4zero, Port: 0,
	// 	},
	// 	Resolver: &net.Resolver{
	// 		PreferGo: true,
	// 		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
	// 			d := net.Dialer{
	// 				Timeout: 30 * time.Second,
	// 			}
	// 			return d.DialContext(ctx, "udp", dnsServer+":"+dnsPort)
	// 		},
	// 	},
	// }
	// serviceDomain := "hello.servic.consul"
	// resolver.Register(&customResolverBuilder{})

	// conn, err := grpc.Dial(serviceDomain, grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name), grpc.WithBlock())
	// defer conn.Close()

	// consul_client, _ := api.NewClient(&api.Config{
	// 	Address: "0.0.0.0:8500",
	// })
	// dis := consul.New(consul_client)
	//endpoint := "discovery://default/hello"
	// conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint(endpoint), grpc.WithDiscovery(dis))

	// conn, err := grpc.Dial("hello.service.consul:9000",
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()),
	// 	grpc.WithResolvers(
	// 		discovery.NewBuilder(
	// 			dis,
	// 			discovery.WithInsecure(true),
	// 		),
	// 	))

	timerCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	conn, err := grpc.DialContext(timerCtx,
		"fyl://hello.service.consul:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithResolvers(
			NewFylBuilder("192.168.10.112", "8600"),
		),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	)
	if err != nil {
		panic(err)
	}
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
