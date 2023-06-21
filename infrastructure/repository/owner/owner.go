package repository

import (
	"github.com/takuya-okada-01/badminist-backend/domain"
	"github.com/takuya-okada-01/badminist-backend/utils"
	"gorm.io/gorm"
)

type ownerRepository struct {
	db *gorm.DB
}

func NewOwnerRepository(db *gorm.DB) domain.IOwnerRepository {
	return &ownerRepository{db: db}
}

func (o *ownerRepository) InsertOwner(owner domain.Owner) error {
	var err error

	result := o.db.Table("owners").Create(owner)
	err = result.Error
	if err != nil {
		return err
	}
	return err
}

func (o *ownerRepository) SelectOwner(criteria domain.OwnerCriteria) (domain.Owner, error) {
	var ownerResult domain.Owner
	mapCriteria := utils.CriteriaToMap(criteria)
	err := o.db.Select("*").Where(mapCriteria).First(&ownerResult).Error
	return ownerResult, err
}

func (o *ownerRepository) SelectOwners(criteria domain.OwnerCriteria) ([]domain.Owner, error) {
	var owners []domain.Owner
	mapCriteria := utils.CriteriaToMap(criteria)
	err := o.db.Select("*").Where(mapCriteria).Find(&owners).Error
	return owners, err
}

func (o *ownerRepository) UpdateOwner(owner *domain.Owner) error {
	err := o.db.Model(&owner).Where(map[string]interface{}{"user_id": owner.UserID, "community_id": owner.CommunityID}).Updates(owner).Error
	return err
}

func (o *ownerRepository) DeleteOwner(userID string, communityID string) error {
	err := o.db.Where(map[string]interface{}{"user_id": userID, "community_id": communityID}).Delete(&domain.Owner{}).Error
	return err
}

func (o *ownerRepository) DeleteOwnerByCommunityID(communityID string) error {
	err := o.db.Where("community_id = ?", communityID).Delete(&domain.Owner{}).Error
	return err
}
