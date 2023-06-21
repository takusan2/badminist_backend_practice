package repository

import (
	"github.com/takuya-okada-01/badminist-backend/domain"
	"github.com/takuya-okada-01/badminist-backend/utils"
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

func (c *communityRepository) SelectCommunity(criteria domain.CommunityCriteria) (domain.Community, error) {
	var community domain.Community
	mapCriteria := utils.CriteriaToMap(criteria)
	err := c.db.Select("*").Where(mapCriteria).First(&community).Error
	return community, err
}

func (c *communityRepository) SelectCommunities(criteria domain.CommunityCriteria) ([]domain.Community, error) {
	var communities []domain.Community
	mapCriteria := utils.CriteriaToMap(criteria)
	err := c.db.Select("*").Where(mapCriteria).Find(&communities).Error
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
