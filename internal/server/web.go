package data

import (
	"context"
	"fmt"
	"helloworld/internal/conf"
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
	return w.srv.ListenAndServe()
}

func (w WebServer) Stop(ctx context.Context) error {
	fmt.Println("stop ")
	w.srv.Shutdown(ctx)
	return nil
}

func NewWebServer() *WebServer {
	fmt.Println(cData.Database.Driver)
	return &WebServer{
		srv:&http.Server{
			Addr:    ":8080",
			Handler: r,
		}
	}
}
