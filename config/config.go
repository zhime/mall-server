package config

type Server struct {
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}
