// api服务启动命令文件
package main

import (
	"rere/routes"
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
	routes.Init(g)

	s := &http.Server{
		Addr: viper.GetString("server.addr"),
		Handler: g,
		ReadTimeout: viper.GetDuration("server.readTimeout") * time.Millisecond,
		WriteTimeout: viper.GetDuration("server.writeTimeout") * time.Millisecond,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}