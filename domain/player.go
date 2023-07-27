package domain

import (
	"time"

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
	NumGames    int       `gorm:"type:int;not null;"`
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
	InsertPlayer(player *Player) (string, error)
	SelectPlayer(playerID string) (Player, error)
	SelectPlayersByCommunityID(communityID string) ([]Player, error)
	SelectAttendPlayers(communityID string) ([]Player, error)
	UpdatePlayer(player *Player) error
	DeletePlayer(id string) error
}
