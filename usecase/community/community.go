package usecase

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/takuya-okada-01/badminist-backend/domain"
)

type communityUseCase struct {
	communityRepository domain.ICommunityRepository
	ownerRepository     domain.IOwnerRepository
}

func NewCommunityUseCase(
	communityRepository domain.ICommunityRepository,
	ownerRepository domain.IOwnerRepository,
) domain.ICommunityUseCase {
	return &communityUseCase{
		communityRepository: communityRepository,
		ownerRepository:     ownerRepository,
	}
}

func (cu *communityUseCase) InsertCommunity(ctx *gin.Context, userID string, community domain.Community) (string, error) {
	id, err := cu.communityRepository.InsertCommunity(&community)
	if err != nil {
		return "", err
	}

	owner := domain.Owner{
		UserID:      userID,
		CommunityID: id,
		Role:        domain.Admin.String(),
	}

	// owner中間テーブルに登録
	err = cu.ownerRepository.InsertOwner(&owner)
	if err != nil {
		return "", err
	}
	return id, err
}

func (cu *communityUseCase) SelectCommunity(ctx *gin.Context, id string) (domain.Community, error) {
	return cu.communityRepository.SelectCommunity(id)
}

func (cu *communityUseCase) SelectCommunitiesByUserID(ctx *gin.Context, userID string) ([]domain.Community, error) {
	return cu.communityRepository.SelectCommunitiesByUserID(userID)
}

func (cu *communityUseCase) UpdateCommunity(ctx *gin.Context, userID string, community *domain.Community) error {
	owner, err := cu.ownerRepository.SelectOwnerByUserIDAndCommunityID(userID, community.ID)
	if err != nil {
		return err
	}
	if owner.Role != domain.Admin.String() && owner.Role != domain.Staff.String() {
		return fmt.Errorf("%s is not authorized to update", owner.Role)
	}

	_, err = cu.communityRepository.SelectCommunity(community.ID)
	if err != nil {
		return err
	}
	return cu.communityRepository.UpdateCommunity(community)
}

func (cu *communityUseCase) DeleteCommunity(ctx *gin.Context, id string) error {
	return cu.communityRepository.DeleteCommunity(id)
}
