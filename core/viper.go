package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/zhime/mall-server/global"
)

func Viper() *viper.Viper {
	v := viper.New()
	/*
		v.SetConfigName("config")
		v.SetConfigType("yaml")
		v.AddConfigPath(".")
		等同于：v.SetConfigFile("config.yaml")
	*/
	v.SetConfigFile("config.yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.MS_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.MS_CONFIG); err != nil {
		panic(err)
	}
	//
	//// root 适配性 根据root位置去找到对应迁移位置,保证root路径有效
	//global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	return v
}
