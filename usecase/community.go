package usecase

import (
	"fmt"

	"github.com/takuya-okada-01/badminist-backend/domain"
)

type CommunityUseCase interface {
	InsertCommunity(userID string, community domain.Community) (string, error)
	SelectCommunityByID(userID string, id string) (domain.CommunityResponse, error)
	SelectCommunitiesByUserID(userID string) ([]domain.CommunityResponse, error)
	UpdateCommunity(userID string, community domain.Community) (domain.CommunityResponse, error)
	DeleteCommunity(userID string, id string) error
}

type communityUseCase struct {
	communityRepository domain.ICommunityRepository
	ownerRepository     domain.IOwnerRepository
}

func NewCommunityUseCase(
	communityRepository domain.ICommunityRepository,
	ownerRepository domain.IOwnerRepository,
) CommunityUseCase {
	return &communityUseCase{
		communityRepository: communityRepository,
		ownerRepository:     ownerRepository,
	}
}

func (cu *communityUseCase) InsertCommunity(userID string, community domain.Community) (string, error) {
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

func (cu *communityUseCase) SelectCommunityByID(userID string, id string) (domain.CommunityResponse, error) {
	_, err := cu.ownerRepository.SelectOwnerByUserIDAndCommunityID(userID, id)
	if err != nil {
		return domain.CommunityResponse{}, err
	}
	community, err := cu.communityRepository.SelectCommunityByID(id)
	if err != nil {
		return domain.CommunityResponse{}, err
	}
	return domain.CommunityResponse{
		ID:          community.ID,
		Name:        community.Name,
		Description: community.Description,
		CreatedAt:   community.CreatedAt,
		UpdatedAt:   community.UpdatedAt,
	}, err
}

func (cu *communityUseCase) SelectCommunitiesByUserID(userID string) ([]domain.CommunityResponse, error) {
	owners, err := cu.ownerRepository.SelectOwnersByUserID(userID)
	if err != nil {
		return nil, err
	}
	communities := make([]domain.CommunityResponse, len(owners))
	for i, owner := range owners {
		community, err := cu.communityRepository.SelectCommunityByID(
			owner.CommunityID,
		)
		if err != nil {
			return nil, err
		}
		communities[i] = domain.CommunityResponse{
			ID:          community.ID,
			Name:        community.Name,
			Description: community.Description,
			CreatedAt:   community.CreatedAt,
			UpdatedAt:   community.UpdatedAt,
		}
	}
	return communities, err
}

func (cu *communityUseCase) UpdateCommunity(userID string, community domain.Community) (domain.CommunityResponse, error) {
	owner, err := cu.ownerRepository.SelectOwnerByUserIDAndCommunityID(userID, community.ID)
	if err != nil {
		return domain.CommunityResponse{}, err
	}
	if owner.Role != domain.Admin.String() && owner.Role != domain.Staff.String() {
		return domain.CommunityResponse{}, fmt.Errorf("%s is not authorized to update", owner.Role)
	}
	_, err = cu.communityRepository.SelectCommunityByID(community.ID)
	if err != nil {
		return domain.CommunityResponse{}, err
	}
	if err := cu.communityRepository.UpdateCommunity(&community); err != nil {
		return domain.CommunityResponse{}, err
	}
	return domain.CommunityResponse{
		ID:          community.ID,
		Name:        community.Name,
		Description: community.Description,
		CreatedAt:   community.CreatedAt,
		UpdatedAt:   community.UpdatedAt,
	}, nil
}

func (cu *communityUseCase) DeleteCommunity(userID string, id string) error {
	owner, err := cu.ownerRepository.SelectOwnerByUserIDAndCommunityID(userID, id)
	if err != nil {
		return err
	}
	if owner.Role != domain.Admin.String() {
		return fmt.Errorf("%s is not authorized to delete", owner.Role)
	}

	return cu.communityRepository.DeleteCommunity(id)
}
