package main

import (
	"flag"
	//"fmt"
	// "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	// "github.com/hashicorp/consul/api"
	"helloworld/internal/data"
	"os"

	"helloworld/internal/conf"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	//服务唯一标识
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string
	//id  采用主机名，不然会替代已有service.ID
	id string
)

func init() {
	//服务唯一标识
	Name = "Demo"
	//id  采用主机名，不然会替代已有service.ID
	id = "ASUS"
	// Version is the version of the compiled software.
	Version = "0.1"
}

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server, demo *data.WebServer, reg *registry.Registry) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
			gs,
			demo,
		),
		kratos.Registrar(reg),
	)
}

func main() {
	flagconf = "/home/fyl/Code/GoProject/kratos_demo/configs/config.yaml"
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Server, bc.Data, bc.Service, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
