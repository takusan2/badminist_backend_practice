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
	Member
)

func (r Role) String() string {
	switch r {
	case Admin:
		return "admin"
	case Member:
		return "member"
	default:
		return "Unknown"
	}
}

type IOwnerRepository interface {
	InsertOwner(owner *Owner) error
	SelectOwner(CommunityID string) (Owner, error)
	SelectOwners(UserID string) ([]Owner, error)
	UpdateOwner(owner *Owner) error
	DeleteOwner(CommunityID string) error
}

type IOwnerUseCase interface {
	InsertOwner(ctx *gin.Context, UserID string, CommunityID string, Role string) error
	SelectOwner(ctx *gin.Context, UserID string, CommunityID string) (Owner, error)
	SelectOwners(ctx *gin.Context, UserID string) ([]Owner, error)
	UpdateOwner(ctx *gin.Context, UserID string, CommunityID string, Role string) error
	DeleteOwner(ctx *gin.Context, UserID string, CommunityID string) error
}
