package auth

import (
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/xyy277/gallery/global"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Init struct {
	useCluster bool
}

func (init *Init) DB(db *gorm.DB) *Init {
	global.G_DB = db
	return init
}

func (init *Init) RedisStandalone(r *redis.Client) *Init {
	global.G_REDIS_MOD = false
	global.G_REDIS_STANDALONE = r
	return init
}

func (init *Init) RedisCluster(rs *redis.ClusterClient) *Init {
	global.G_REDIS_MOD = true
	global.G_REDIS_CLUSTER = rs
	return init
}

func (init *Init) Viper(vp *viper.Viper) *Init {
	global.G_VP = vp
	return init
}

func (init *Init) Zap(zap *zap.Logger) *Init {
	global.G_LOG = zap
	return init
}

func (init *Init) BlackCache(cache local_cache.Cache) *Init {
	global.BlackCache = cache
	return init
}
