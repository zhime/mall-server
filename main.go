package main

import (
	"github.com/zhime/mall-server/core"
	"github.com/zhime/mall-server/global"
)

func main() {
	global.MS_VP = core.Viper()

}
