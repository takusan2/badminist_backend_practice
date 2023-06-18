package domain

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Player struct {
	ID          string    `gorm:"type:varchar(36);primary_key;"`
	Name        string    `gorm:"type:varchar(255);not null;"`
	Sex         bool      `gorm:"type:boolean;not null;"` // sex 0: 女性, 1: 男性
	Age         int       `gorm:"type:int;not null;"`
	Level       int       `gorm:"type:int;not null;"`
	Attendance  bool      `gorm:"type:boolean;not null;"`
	CommunityID string    `gorm:"type:varchar(36);not null;"`
	Community   Community `gorm:"foreignkey:CommunityID"`
	CreatedAt   time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;"`
	UpdatedAt   time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;"`
}

func (p *Player) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New().String()
	return nil
}

type IPlayerRepository interface {
	InsertPlayer(communityID string, player *Player) (string, error)
	SelectPlayer(id string) (Player, error)
	SelectPlayersByCommunityID(communityID string) ([]Player, error)
	UpdatePlayer(player *Player) error
	DeletePlayer(id string) error
}

type IPlayerUseCase interface {
	InsertPlayer(ctx *gin.Context, communityID string, player *Player) (string, error)
	SelectPlayer(ctx *gin.Context, id string) (Player, error)
	SelectPlayersByCommunityID(ctx *gin.Context, communityID string) ([]Player, error)
	UpdatePlayer(ctx *gin.Context, player *Player) error
	DeletePlayer(ctx *gin.Context, id string) error
}
