package repository

import (
	"github.com/takuya-okada-01/badminist-backend/domain"
	"gorm.io/gorm"
)

type ownerRepository struct {
	db *gorm.DB
}

func NewOwnerRepository(db *gorm.DB) domain.IOwnerRepository {
	return &ownerRepository{db: db}
}

func (o *ownerRepository) InsertOwner(owner *domain.Owner) error {
	var err error

	result := o.db.Table("owners").Create(owner)
	err = result.Error
	if err != nil {
		return err
	}
	return err
}

func (o *ownerRepository) SelectOwner(owner domain.Owner) (domain.Owner, error) {
	var ownerResult domain.Owner
	err := o.db.Select("*").Where(map[string]interface{}{"user_id": owner.UserID, "community_id": owner.CommunityID}).First(&ownerResult).Error
	return ownerResult, err
}

func (o *ownerRepository) SelectOwnerByUserIDAndCommunityID(userID, communityId string) (domain.Owner, error) {
	var owner domain.Owner
	err := o.db.Select("*").Where(map[string]interface{}{"user_id": userID, "community_id": communityId}).First(&owner).Error
	return owner, err
}

func (o *ownerRepository) SelectOwnersByCommunityID(communityId string) ([]domain.Owner, error) {
	var owners []domain.Owner
	err := o.db.Select("*").Where("community_id = ?", communityId).Find(&owners).Error
	return owners, err
}

func (o *ownerRepository) UpdateOwner(owner *domain.Owner) error {
	err := o.db.Model(&owner).Where(map[string]interface{}{"user_id": owner.UserID, "community_id": owner.CommunityID}).Updates(owner).Error
	return err
}

func (o *ownerRepository) DeleteOwner(owner domain.Owner) error {
	err := o.db.Where(map[string]interface{}{"user_id": owner.UserID, "community_id": owner.CommunityID}).Delete(&domain.Owner{}).Error
	return err
}

func (o *ownerRepository) DeleteOwnerByCommunityID(communityID string) error {
	err := o.db.Where("community_id = ?", communityID).Delete(&domain.Owner{}).Error
	return err
}
