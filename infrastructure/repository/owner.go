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

func (o *ownerRepository) InsertOwner(owner domain.Owner) error {
	var err error

	result := o.db.Table("owners").Create(owner)
	err = result.Error
	if err != nil {
		return err
	}
	return err
}

func (o *ownerRepository) SelectOwnerByUserIDAndCommunityID(userID string, communityID string) (domain.Owner, error) {
	var ownerResult domain.Owner
	err := o.db.Select("*").Where(map[string]any{
		"user_id":      userID,
		"community_id": communityID,
	}).First(&ownerResult).Error
	return ownerResult, err
}

func (o *ownerRepository) SelectOwnersByCommunityID(communityID string) ([]domain.Owner, error) {
	var owners []domain.Owner
	err := o.db.Select("*").Where("community_id = ?", communityID).Find(&owners).Error
	return owners, err
}

func (o *ownerRepository) SelectOwnersByUserID(userID string) ([]domain.Owner, error) {
	var owners []domain.Owner
	err := o.db.Select("*").Where("user_id = ?", userID).Find(&owners).Error
	return owners, err
}

func (o *ownerRepository) UpdateOwner(owner *domain.Owner) error {
	err := o.db.Model(&owner).Where(map[string]any{
		"user_id":      owner.UserID,
		"community_id": owner.CommunityID,
	}).Updates(owner).Error
	return err
}

func (o *ownerRepository) DeleteOwner(userID string, communityID string) error {
	err := o.db.Where(map[string]any{
		"user_id":      userID,
		"community_id": communityID,
	}).Delete(&domain.Owner{}).Error
	return err
}

func (o *ownerRepository) DeleteOwnerByCommunityID(communityID string) error {
	err := o.db.Where("community_id = ?", communityID).Delete(&domain.Owner{}).Error
	return err
}
