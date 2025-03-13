package main

import (
	"github.com/zhime/mall-server/core"
	"github.com/zhime/mall-server/global"
	"github.com/zhime/mall-server/initialize"
)

func main() {
	global.MS_VP = core.Viper()
	//global.MS_LOG = core.Zap()
	//global.MS_LOG.Info("TEST")
	r := initialize.Routers()
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
