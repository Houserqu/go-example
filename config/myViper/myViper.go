package myViper

import (
	"fmt"

	"github.com/spf13/viper"
)

func RunViper() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.AutomaticEnv() // 加载环境变量

	err := viper.ReadInConfig() // 搜索并读取配置文件
	if err != nil {             // 处理错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	viper.SetDefault("NAME", "hahaha")

	fmt.Println(viper.GetStringMap("global"))
	fmt.Println(viper.Get("APP"))
}
