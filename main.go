package main

import (
	"fmt"
	"github.com/zhime/mall-server/core"
	"github.com/zhime/mall-server/global"
)

func main() {
	global.MS_VP = core.Viper()
	fmt.Println(global.MS_CONFIG)
}
