// api服务启动命令文件
package main

import (
	"time"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"rere/bootstrap"
)

func main() {
	bootstrap.Init()

	run()
}

func run()  {
	g := gin.Default()
	g.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})


	s := &http.Server{
		Addr: viper.GetString("server.addr"),
		Handler: g,
		ReadTimeout: viper.GetDuration("server.readTimeout") * time.Millisecond,
		WriteTimeout: viper.GetDuration("server.writeTimeout") * time.Millisecond,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}