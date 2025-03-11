package core

import (
	"fmt"
	"github.com/zhime/mall-server/core/internal"
	"github.com/zhime/mall-server/global"
	"github.com/zhime/mall-server/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.MS_CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.MS_CONFIG.Zap.Director)
		_ = os.Mkdir(global.MS_CONFIG.Zap.Director, os.ModePerm)
	}
	levels := global.MS_CONFIG.Zap.Levels()
	length := len(levels)
	cores := make([]zapcore.Core, 0, length)
	for i := 0; i < length; i++ {
		core := internal.NewZapCore(levels[i])
		cores = append(cores, core)
	}
	logger = zap.New(zapcore.NewTee(cores...))
	if global.MS_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
