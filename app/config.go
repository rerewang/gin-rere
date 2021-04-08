package app

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig()  {
	// 读取配置文件
	cfgFile := flag.String("c", "config/dev/conf.toml", "config file")
	flag.Parse()

	cfg := viper.GetViper()
	cfg.SetConfigFile(*cfgFile)
	err := cfg.ReadInConfig()

	if err != nil {
		msg := fmt.Sprintf("read config file error file=%s", cfgFile)
		panic(msg)
	}

	G.Cfg = cfg
}
