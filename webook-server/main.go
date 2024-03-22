package main

import (
	_ "github.com/spf13/viper/remote"
	"time"
	"webook-server/pkg/snowflake"
)

const (
	//	rate limit
	window = time.Second
	rate   = 100
)

func main() {
	snowflake.Init("2023-01-01", 1)
	r := InitWebServer(window, rate)
	////initViperV2()
	_ = r.Run(":8080")
}

//func initViper() {
//	viper.SetConfigName("dev")
//	viper.SetConfigType("yaml")
//	viper.AddConfigPath("./config")
//	if err := viper.ReadInConfig(); err != nil {
//		panic(err)
//	} // 加载到内存
//}
//
//func initViperV1() {
//	cFile := pflag.String("config", "./dev.yaml", "指定配置文件路径")
//	pflag.Parse()
//	viper.SetConfigFile(*cFile)
//}
//func initViperV2() {
//	viper.SetConfigType("yaml")
//	if err := viper.AddRemoteProvider("etcd3", "127.0.0.1:23790", "/webook"); err != nil {
//		panic(err)
//	}
//	if err := viper.ReadRemoteConfig(); err != nil {
//		panic(err)
//	}
//}
//
//func initViperV3() {
//	viper.SetConfigFile(confFile) // 指定配置文件路径
//	// 读取配置信息
//	if err := viper.ReadInConfig(); err != nil { // 读取配置信息失败
//		panic(fmt.Errorf("Fatal error config file: %s \n", err))
//	}
//	// 将读取的配置信息保存至全局变量Conf
//	if err := viper.Unmarshal(Conf); err != nil {
//		panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
//	}
//	// 监控配置文件变化
//	viper.WatchConfig()
//	// 注意！！！配置文件发生变化后要同步到全局变量Conf
//	viper.OnConfigChange(func(in fsnotify.Event) {
//		fmt.Println("夭寿啦~配置文件被人修改啦...")
//		if err := viper.Unmarshal(Conf); err != nil {
//			panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
//		}
//	})
//}
