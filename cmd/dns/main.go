package main

import (
	"context"
	"fmt"
	"net"
	"time"
)

func main() {

	dnsServer := "0.0.0.0"
	dnsPort := "8600"
	serverHost := "hello.service.consul"
	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: 30 * time.Second,
			}
			return d.DialContext(ctx, "udp", dnsServer+":"+dnsPort)
		},
	}
	ips, err := resolver.LookupIP(context.Background(), "ip4", serverHost)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(ips)
	}

}
