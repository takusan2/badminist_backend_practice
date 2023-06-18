package repository

import (
	"github.com/takuya-okada-01/badminist-backend/domain"
	"gorm.io/gorm"
)

type communityRepository struct {
	db *gorm.DB
}

func NewCommunityRepository(db *gorm.DB) domain.ICommunityRepository {
	return &communityRepository{db: db}
}

func (c *communityRepository) InsertCommunity(community *domain.Community) (string, error) {
	var err error
	result := c.db.Table("communities").Create(community)
	err = result.Error
	if err != nil {
		return "", err
	}
	return community.ID, err
}

func (c *communityRepository) SelectCommunity(id string) (domain.Community, error) {
	var community domain.Community
	err := c.db.Select("*").Where("id = ?", id).First(&community).Error
	return community, err
}

func (c *communityRepository) SelectCommunities(UserID string) ([]domain.Community, error) {
	var communities []domain.Community
	err := c.db.Select("*").Where("user_id = ?", UserID).Find(&communities).Error
	return communities, err
}

func (c *communityRepository) UpdateCommunity(community *domain.Community) error {
	err := c.db.Model(&community).Where("id = ?", community.ID).Updates(community).Error
	return err
}

func (c *communityRepository) DeleteCommunity(id string) error {
	err := c.db.Where("id = ?", id).Delete(&domain.Community{}).Error
	return err
}
