package main

import (
	"github.com/zhime/mall-server/core"
	"github.com/zhime/mall-server/global"
)

func main() {
	global.MS_VP = core.Viper()
	global.MS_LOG = core.Zap()
	global.MS_LOG.Info("TEST")
}
