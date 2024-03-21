package main

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"webook-server/pkg/snowflake"
)

func main() {
	snowflake.Init("2023-01-01", 1)
	r := InitWebServer()
	_ = r.Run(":8080")
}

func initViper() {
	viper.SetConfigName("dev")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	} // 加载到内存
}

func initViperV1() {
	cFile := pflag.String("config", "./dev.yaml", "指定配置文件路径")
	pflag.Parse()
	viper.SetConfigFile(*cFile)
}
