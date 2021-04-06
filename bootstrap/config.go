package bootstrap

import (
	"fmt"
	"github.com/spf13/viper"
)

func initConfig()  {
	cfgFile := "config/dev/conf.toml"
	cfg := viper.GetViper()
	cfg.SetConfigFile(cfgFile)
	err := cfg.ReadInConfig()

	if err != nil {
		msg := fmt.Sprintf("read config file error file=%s", cfgFile)
		panic(msg)
	}
}