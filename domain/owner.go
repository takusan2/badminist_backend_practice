package domain

import (
	"github.com/gin-gonic/gin"
)

type Owner struct {
	UserID      string    `gorm:"primary_key;"`
	User        User      `gorm:"foreignKey:UserID;references:ID"`
	CommunityID string    `gorm:"primary_key;"`
	Community   Community `gorm:"foreignKey:CommunityID;references:ID"`
	Role        string    `gorm:"type:varchar(255);not null;"`
}

type OwnerCriteria struct {
	UserID               string
	UserIDIsNotNull      bool
	CommunityID          string
	CommunityIDIsNotNull bool
	Role                 string
	RoleIsNotNull        bool
}

type IOwnerRepository interface {
	InsertOwner(owner Owner) error
	SelectOwner(criteria OwnerCriteria) (Owner, error)
	SelectOwners(criteria OwnerCriteria) ([]Owner, error)
	UpdateOwner(owner *Owner) error
	DeleteOwner(userID string, communityID string) error
}

type IOwnerUseCase interface {
	InsertOwner(ctx *gin.Context, owner Owner) error
	SelectOwnersByCommunityID(ctx *gin.Context, communityID string) ([]Owner, error)
	SelectOwnerByUserIDAndCommunityID(ctx *gin.Context, userID string, communityID string) (Owner, error)
	UpdateOwner(ctx *gin.Context, userID string, owner *Owner) error
	DeleteOwner(ctx *gin.Context, userID string, delUserID string, delCommunityID string) error
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
		return "unknown"
	}
}
