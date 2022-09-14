package global

import (
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"gallery/gsafety.com/config"
)

var (
	G_DB                  *gorm.DB
	G_REDIS_STANDALONE    *redis.Client
	G_REDIS_CLUSTER       *redis.ClusterClient
	G_REDIS_MOD           bool
	G_CONFIG              config.Server
	G_VP                  *viper.Viper
	G_LOG                 *zap.Logger
	G_Concurrency_Control = &singleflight.Group{}

	BlackCache local_cache.Cache
	lock       sync.RWMutex
)

type Init struct {
	useCluster bool
}

func (init *Init) DB(db *gorm.DB) *Init {
	G_DB = db
	return init
}

func (init *Init) RedisStandalone(r *redis.Client) *Init {
	G_REDIS_MOD = false
	G_REDIS_STANDALONE = r
	return init
}

func (init *Init) RedisCluster(rs *redis.ClusterClient) *Init {
	G_REDIS_MOD = true
	G_REDIS_CLUSTER = rs
	return init
}

func (init *Init) Viper(vp *viper.Viper) *Init {
	G_VP = vp
	return init
}

func (init *Init) Zap(zap *zap.Logger) *Init {
	G_LOG = zap
	return init
}
