package config

type Server struct {
	JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	Casbin Casbin `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
}
