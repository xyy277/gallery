package config

type System struct {
	Application   string `mapstructure:"application" json:"application" yaml:"application"`          // application
	Host          string `mapstructure:"host" json:"host" yaml:"host"`                               // host
	Port          string `mapstructure:"port" json:"port" yaml:"port"`                               // port
	RootPath      string `mapstructure:"root-path" json:"root-path" yaml:"root-path"`                // root-path
	DbType        string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`                      // 数据库类型
	OssType       string `mapstructure:"oss-type" json:"oss-type" yaml:"oss-type"`                   // Oss类型
	Mode          string `mapstructure:"mode" json:"mode" yaml:"mode"`                               // 环境值
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"use-multipoint" yaml:"use-multipoint"` // 多点登录拦截
	UseRedis      bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`                // 使用redis
	UseDatabase   bool   `mapstructure:"use-database" json:"use-database" yaml:"use-database"`       // 使用数据库
	AutoMigrate   bool   `mapstructure:"auto-migrate" json:"auto-migrate" yaml:"auto-migrate"`       // 自动建表
	TimeZone      string `mapstructure:"time-zone" json:"time-zone" yaml:"time-zone"`                // 时区
	LimitCountIP  int    `mapstructure:"iplimit-count" json:"iplimit-count" yaml:"iplimit-count"`
	LimitTimeIP   int    `mapstructure:"iplimit-time" json:"iplimit-time" yaml:"iplimit-time"`
}
