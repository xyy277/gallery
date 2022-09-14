package config

type Server struct {
	JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Casbin Casbin `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
}
