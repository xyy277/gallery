package auth

import (
	"gallery/gsafety.com/global"

	"go.uber.org/zap"
)

// 权限接口
type JWT interface {
	JsonInBlacklist(jwtList JwtBlacklist) (err error)
	IsBlacklist(jwt string) bool
	GetRedisJWT(userName string) (redisJWT string, err error)
	SetRedisJWT(jwt string, userName string) (err error)
}

// @author: [zhoujiajun](zhoujiajun@gsafety.com)
// 权限实现
func NewJWT() JWT {

	return &JwtService{}
}

func LoadAll() {
	var data []string
	err := global.G_DB.Model(&JwtBlacklist{}).Select("jwt").Find(&data).Error
	if err != nil {
		global.G_LOG.Error("加载数据库jwt黑名单失败!", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		global.BlackCache.SetDefault(data[i], struct{}{})
	} // jwt黑名单 加入 BlackCache 中
}
