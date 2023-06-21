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
	err = cu.ownerRepository.InsertOwner(owner)
	if err != nil {
		return "", err
	}
	return id, err
}

func (cu *communityUseCase) SelectCommunity(ctx *gin.Context, id string) (domain.Community, error) {
	return cu.communityRepository.SelectCommunity(domain.CommunityCriteria{
		ID:          id,
		IDIsNotNull: true,
	})
}

func (cu *communityUseCase) SelectCommunitiesByUserID(ctx *gin.Context, userID string) ([]domain.Community, error) {
	owners, err := cu.ownerRepository.SelectOwners(
		domain.OwnerCriteria{
			UserID:          userID,
			UserIDIsNotNull: true,
		})
	if err != nil {
		return nil, err
	}
	communities := make([]domain.Community, len(owners))
	for i, owner := range owners {
		community, err := cu.communityRepository.SelectCommunity(domain.CommunityCriteria{
			ID:          owner.CommunityID,
			IDIsNotNull: true,
		})
		if err != nil {
			return nil, err
		}
		communities[i] = community
	}
	return communities, err
}

func (cu *communityUseCase) UpdateCommunity(ctx *gin.Context, userID string, community *domain.Community) error {
	owner, err := cu.ownerRepository.SelectOwner(domain.OwnerCriteria{
		UserID:               userID,
		UserIDIsNotNull:      true,
		CommunityID:          community.ID,
		CommunityIDIsNotNull: true,
	})
	fmt.Print("OK1\n")
	if err != nil {
		return err
	}
	fmt.Print("OK2\n")
	if owner.Role != domain.Admin.String() && owner.Role != domain.Staff.String() {
		return fmt.Errorf("%s is not authorized to update", owner.Role)
	}
	fmt.Print("OK3\n")

	_, err = cu.communityRepository.SelectCommunity(domain.CommunityCriteria{
		ID:          community.ID,
		IDIsNotNull: true,
	})
	fmt.Print("OK4\n")
	if err != nil {
		return err
	}
	return cu.communityRepository.UpdateCommunity(community)
}

func (cu *communityUseCase) DeleteCommunity(ctx *gin.Context, id string) error {
	return cu.communityRepository.DeleteCommunity(id)
}
