package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebServer struct {
	srv *http.Server
}

func (w WebServer) Start(ctx context.Context) error {
	fmt.Println("web server start")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		//输出json结果给调用方
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r.Run(":8081")
}

func (w WebServer) Stop(ctx context.Context) error {
	fmt.Println("stop ")
	return nil
}

func NewWebServer() *WebServer {
	return &WebServer{}
}
