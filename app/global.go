package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

var G *Global

// G 表示全局变量
type Global struct {
	Log *logrus.Logger
	Cfg *viper.Viper
	Env *Env
	Now time.Time
}

func InitGlobalVar(g *gin.Engine) {
	// 初始化全局结构体
	G = &Global{
		Now: time.Now().Local(),
	}

	//初始化配置
	InitConfig()

	//初始化ENV
	InitEnv()

	//初始化Log
	InitLog()
}