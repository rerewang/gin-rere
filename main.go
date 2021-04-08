// api服务启动命令文件
package main

import (
	"net/http"
	"rere/internal/api"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"rere/app"
)

func main() {
	g := gin.Default()

	app.InitGlobalVar(g)

	run(g)
}

func run(g *gin.Engine)  {
	s := &http.Server{
		Addr: viper.GetString("server.addr"),
		Handler: g,
		ReadTimeout: viper.GetDuration("server.readTimeout") * time.Millisecond,
		WriteTimeout: viper.GetDuration("server.writeTimeout") * time.Millisecond,
		MaxHeaderBytes: 1 << 20,
	}

	if err := api.Init(g); err != nil {
		app.G.Log.Fatal(err)
	}

	s.ListenAndServe()
}