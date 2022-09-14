package auth

import (
	"github.com/xyy277/gallery/global"

	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
)

type JwtBlacklist struct {
	global.GS_BASE_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}

// Custom claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.StandardClaims
}

type BaseClaims struct {
	UUID        uuid.UUID
	ID          uint
	Username    string
	NickName    string
	AuthorityId string
}
