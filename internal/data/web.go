package data

import (
	"context"
	"fmt"
	"helloworld/internal/conf"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebServer struct {
}

var srv *http.Server

func (w WebServer) Start(ctx context.Context) error {
	fmt.Println("wev server start")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		//输出json结果给调用方
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	srv = &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	return srv.ListenAndServe()
}

func (w WebServer) Stop(ctx context.Context) error {
	fmt.Println("stop ")
	srv.Shutdown(ctx)
	return nil
}

func NewWebServer(cData *conf.Data) *WebServer {
	fmt.Println(cData.Database.Driver)
	return &WebServer{}
}
