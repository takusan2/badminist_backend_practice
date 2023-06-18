package domain

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Player struct {
	ID   string `gorm:"type:varchar(36);primary_key;"`
	Name string `gorm:"type:varchar(255);not null;"`
	// sex 0: 女性, 1: 男性
	Sex bool `gorm:"type:boolean;not null;"`
	// level max: 10
	Level      int       `gorm:"type:int;not null;"`
	Attendance bool      `gorm:"type:boolean;not null;"`
	CreatedAt  time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;"`
	UpdatedAt  time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;"`
}

func (p *Player) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New().String()
	return nil
}

type IPlayerRepository interface {
	InsertPlayer(userID string, player *Player) (string, error)
	SelectPlayer(id string) (Player, error)
	SelectPlayers(communityID string) ([]Player, error)
	UpdatePlayer(player *Player) error
	DeletePlayer(id string) error
}

type IPlayerUseCase interface {
	InsertPlayer(ctx *gin.Context, userID string, player *Player) (string, error)
	SelectPlayer(ctx *gin.Context, id string) (Player, error)
	SelectPlayers(ctx *gin.Context, communityID string) ([]Player, error)
	UpdatePlayer(ctx *gin.Context, player *Player) error
	DeletePlayer(ctx *gin.Context, id string) error
}
