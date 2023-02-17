package config

import "github.com/spf13/viper"

var C *Config

func init() {
	var configViperConfig = viper.New()
	configViperConfig.SetConfigName("config")
	configViperConfig.SetConfigType("yaml")
	configViperConfig.AddConfigPath("common/config/")
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
