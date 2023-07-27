package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
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

type CommunityResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type ICommunityRepository interface {
	InsertCommunity(community *Community) (string, error)
	SelectCommunityByID(id string) (Community, error)
	UpdateCommunity(community *Community) error
	DeleteCommunity(id string) error
}
