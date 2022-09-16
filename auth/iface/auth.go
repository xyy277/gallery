package iface

import (
	"github.com/xyy277/gallery/global/response"
)

// star 请求 luna的接口进行鉴权与授权
type AUTH interface {
	// 通过token请求鉴权
	RequestToken(token string) (bool, response.Response)
	// 通过request中的path和method请求验证策略
	RequestEnforce(path string, method string) bool
}

func NewStarAUTH() AUTH {
	// Request impl
	return &star.authentication{}
}
