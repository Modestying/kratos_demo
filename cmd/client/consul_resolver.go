package main

import (
	"context"
	"fmt"
	"net"
	"strings"
	"time"

	"google.golang.org/grpc/resolver"
)

func init() {
	resolver.Register(&FylBuilder{})
}

const Scheme = "fyl"

type FylBuilder struct {
	dnsServer string
	dnsPort   string
}

var _ resolver.Builder = (*FylBuilder)(nil)

func NewFylBuilder(addr, port string) resolver.Builder {
	return &FylBuilder{
		dnsServer: addr,
		dnsPort:   port,
	}
}
func (b *FylBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &FResolver{
		dnsServer: b.dnsServer,
		dnsPort:   b.dnsPort,
		target:    target.URL.Host,
		cc:        cc,
		resolver: &net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				d := net.Dialer{
					Timeout: 30 * time.Second,
				}
				return d.DialContext(ctx, "udp", b.dnsServer+":"+b.dnsPort)
			},
		},
	}
	r.ResolveNow(resolver.ResolveNowOptions{})
	return r, nil
}
func (b *FylBuilder) Scheme() string {
	return Scheme
}

type FResolver struct {
	dnsServer string
	dnsPort   string
	target    string
	resolver  *net.Resolver
	cc        resolver.ClientConn
}

var _ resolver.Resolver = (*FResolver)(nil)

func (f *FResolver) ResolveNow(opts resolver.ResolveNowOptions) {
	fmt.Println(f.target)
	tagetServerInfo := strings.Split(f.target, ":")
	if len(tagetServerInfo) != 2 {
		panic("url post format error,retry with correct format,like: fyl://hello.service.consul:9000")
	}
	ips, err := f.resolver.LookupIP(context.Background(), "ip4", tagetServerInfo[0])
	if err != nil {
		panic(err)
	}
	fmt.Println(ips)
	addrs := make([]resolver.Address, len(ips))
	for i, s := range ips {
		addrs[i] = resolver.Address{Addr: s.String() + ":" + tagetServerInfo[1]}
	}
	f.cc.UpdateState(resolver.State{
		Addresses: addrs,
	})
}

// Close closes the resolver.
func (f *FResolver) Close() {
}
