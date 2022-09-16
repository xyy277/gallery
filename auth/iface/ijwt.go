package iface

import "github.com/xyy277/gallery/auth/luna"

// 权限接口
type JWT interface {
	JsonInBlacklist(jwtList luna.JwtBlacklist) (err error)
	IsBlacklist(jwt string) bool
	GetCacheJWT(userName string) (redisJWT string, err error)
	SetCacheJWT(jwt string, userName string) (err error)
}

// @author: [zhoujiajun](zhoujiajun@gsafety.com)
// 权限实现
func NewJWT() JWT {

	// 不同实现
	return &luna.JwtService{}
}
