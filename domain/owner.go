package domain

import (
	"github.com/gin-gonic/gin"
)

type Owner struct {
	UserID      string
	User        User `gorm:"foreignKey:UserID;references:ID"`
	CommunityID string
	Community   Community `gorm:"foreignKey:CommunityID;references:ID"`
	Role        string    `gorm:"type:varchar(255);not null;"`
}

type Role int

const (
	Admin Role = iota
	Staff
	Member
)

func (r Role) String() string {
	switch r {
	case Admin:
		return "admin"
	case Staff:
		return "staff"
	case Member:
		return "member"
	default:
		return "Unknown"
	}
}

type IOwnerRepository interface {
	InsertOwner(owner *Owner) error
	SelectOwner(owner Owner) (Owner, error)
	SelectOwnersByCommunityID(communityID string) ([]Owner, error)
	SelectOwnerByUserIDAndCommunityID(userID string, communityID string) (Owner, error)
	UpdateOwner(owner *Owner) error
	DeleteOwner(owner Owner) error
}

type IOwnerUseCase interface {
	InsertOwner(ctx *gin.Context, userID string, communityID string, Role string) error
	SelectOwner(ctx *gin.Context, userID string, communityID string) (Owner, error)
	SelectOwnersByCommunityID(ctx *gin.Context, UserID string) ([]Owner, error)
	UpdateOwner(ctx *gin.Context, userID string, communityID string, Role string) error
	DeleteOwner(ctx *gin.Context, userID string, communityID string) error
}
