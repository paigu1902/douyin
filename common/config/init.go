package config

import (
	"github.com/spf13/viper"
	"path"
	"runtime"
)

var C *Config

func init() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("配置初始化失败")
	}
	var configViperConfig = viper.New()
	configViperConfig.SetConfigName("config")
	configViperConfig.SetConfigType("yaml")
	configViperConfig.AddConfigPath(path.Dir(filename))

	//读取配置文件内容
	if err := configViperConfig.ReadInConfig(); err != nil {
		panic(err)
	}
	var c Config
	if err := configViperConfig.Unmarshal(&c); err != nil {
		panic(err)
	}
	C = &c
}
