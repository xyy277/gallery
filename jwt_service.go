package gallery

import "gallery/auth"

// 权限接口
type JWT interface {
	JsonInBlacklist(jwtList auth.JwtBlacklist) (err error)
	IsBlacklist(jwt string) bool
	GetRedisJWT(userName string) (redisJWT string, err error)
	SetRedisJWT(jwt string, userName string) (err error)
}

// @author: [zhoujiajun](zhoujiajun@gsafety.com)
// 权限实现
func NewJWT() JWT {

	return &auth.JwtService{}
}
