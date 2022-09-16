package star

import (
	uuid "github.com/satori/go.uuid"
	"github.com/xyy277/gallery/auth/luna"
	"github.com/xyy277/gallery/auth/star"
)

type CLAIM interface {
	GetClaims(token string) (*luna.CustomClaims, error)
	GetUserID(token string) uint
	GetUserUUID(token string) uuid.UUID
	GetUserAuthorityId(token string) string
}

func NewStarCLAMI() CLAIM {

	return &star.Claimant{}
}
