package domain

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Community struct {
	ID          string    `gorm:"type:varchar(36);primary_key;"`
	Name        string    `gorm:"type:varchar(255);not null;"`
	Description string    `gorm:"type:text;"`
	CreatedAt   time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;"`
	UpdatedAt   time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;"`
}

func (c *Community) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()
	return nil
}

type ICommunityRepository interface {
	InsertCommunity(UserID string, community *Community) (string, error)
	SelectCommunity(id string) (Community, error)
	SelectCommunities(UserID string) ([]Community, error)
	UpdateCommunity(UserID string, community *Community) error
	DeleteCommunity(id string) error
}

type ICommunityUseCase interface {
	InsertCommunity(ctx *gin.Context, UserID string, community Community) (string, error)
	SelectCommunity(ctx *gin.Context, id string) (Community, error)
	SelectCommunities(ctx *gin.Context, UserID string) ([]Community, error)
	UpdateCommunity(ctx *gin.Context, UserID string, community *Community) error
	DeleteCommunity(ctx *gin.Context, id string) error
}
