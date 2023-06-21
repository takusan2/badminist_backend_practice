package usecase

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/takuya-okada-01/badminist-backend/domain"
)

type ownerUseCase struct {
	ownerRepository domain.IOwnerRepository
}

func NewOwnerUseCase(ownerRepository domain.IOwnerRepository) domain.IOwnerUseCase {
	return &ownerUseCase{
		ownerRepository: ownerRepository,
	}
}

func (ouc *ownerUseCase) InsertOwner(ctx *gin.Context, owner domain.Owner) error {
	return ouc.ownerRepository.InsertOwner(owner)
}

func (ouc *ownerUseCase) SelectOwnersByCommunityID(ctx *gin.Context, communityID string) ([]domain.Owner, error) {
	return ouc.ownerRepository.SelectOwners(domain.OwnerCriteria{
		CommunityID:          communityID,
		CommunityIDIsNotNull: true,
	})
}

func (ouc *ownerUseCase) SelectOwnerByUserIDAndCommunityID(ctx *gin.Context, userID string, communityID string) (domain.Owner, error) {
	return ouc.ownerRepository.SelectOwner(
		domain.OwnerCriteria{
			UserID:               userID,
			UserIDIsNotNull:      true,
			CommunityID:          communityID,
			CommunityIDIsNotNull: true,
		})
}

func (ouc *ownerUseCase) UpdateOwner(ctx *gin.Context, userID string, owner *domain.Owner) error {
	fetchOwner, err := ouc.ownerRepository.SelectOwner(
		domain.OwnerCriteria{
			UserID:               userID,
			UserIDIsNotNull:      true,
			CommunityID:          owner.CommunityID,
			CommunityIDIsNotNull: true,
		})
	if err != nil {
		return err
	}
	if fetchOwner.Role != domain.Admin.String() {
		return fmt.Errorf("Only the admin is authorized to update")
	}

	return ouc.ownerRepository.UpdateOwner(owner)
}

func (ouc *ownerUseCase) DeleteOwner(ctx *gin.Context, userID string, delUserID string, delCommunityID string) error {
	fetchOwner, err := ouc.ownerRepository.SelectOwner(
		domain.OwnerCriteria{
			UserID:               userID,
			UserIDIsNotNull:      true,
			CommunityID:          delCommunityID,
			CommunityIDIsNotNull: true,
		})
	if err != nil {
		return err
	}
	if fetchOwner.Role != domain.Admin.String() {
		return fmt.Errorf("Only the admin is authorized to update")
	}

	return ouc.ownerRepository.DeleteOwner(delUserID, delCommunityID)
}
