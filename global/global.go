package global

import (
	"github.com/spf13/viper"
	"github.com/zhime/mall-server/config"
	"go.uber.org/zap"
)

var (
	MS_VP     *viper.Viper
	MS_CONFIG *config.Server
	MS_LOG    *zap.Logger
)
