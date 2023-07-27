package usecase

import (
	"fmt"

	"github.com/takuya-okada-01/badminist-backend/domain"
)

type OwnerUseCase interface {
	InsertOwner(userID string, owner domain.Owner) error
	SelectOwnersByCommunityID(userID, communityID string) ([]domain.Owner, error)
	UpdateOwner(userID string, owner *domain.Owner) error
	DeleteOwner(userID string, delUserID string, delCommunityID string) error
}

type ownerUseCase struct {
	or domain.IOwnerRepository
	cr domain.ICommunityRepository
}

func NewOwnerUseCase(
	or domain.IOwnerRepository,
	cr domain.ICommunityRepository,
) OwnerUseCase {
	return &ownerUseCase{
		or: or,
		cr: cr,
	}
}

func (ouc *ownerUseCase) InsertOwner(userID string, owner domain.Owner) error {
	community, err := ouc.cr.SelectCommunityByID(owner.CommunityID)
	if err != nil {
		return err
	}
	if community.ID == "" {
		return fmt.Errorf("Not found the community")
	}
	fetchOwner, err := ouc.or.SelectOwnerByUserIDAndCommunityID(
		userID,
		owner.CommunityID,
	)
	if err != nil {
		return err
	}
	if fetchOwner.Role != domain.Admin.String() {
		return fmt.Errorf("Only the admin is authorized to insert")
	}
	return ouc.or.InsertOwner(owner)
}

func (ouc *ownerUseCase) SelectOwnersByCommunityID(
	userID string,
	communityID string,
) ([]domain.Owner, error) {
	_, err := ouc.or.SelectOwnerByUserIDAndCommunityID(
		userID,
		communityID,
	)
	if err != nil {
		return nil, err
	}
	return ouc.or.SelectOwnersByCommunityID(communityID)
}

func (ouc *ownerUseCase) UpdateOwner(userID string, owner *domain.Owner) error {
	fetchOwner, err := ouc.or.SelectOwnerByUserIDAndCommunityID(
		userID,
		owner.CommunityID,
	)
	if err != nil {
		return err
	}
	if fetchOwner.Role != domain.Admin.String() {
		return fmt.Errorf("Only the admin is authorized to update")
	}

	return ouc.or.UpdateOwner(owner)
}

func (ouc *ownerUseCase) DeleteOwner(userID string, delUserID string, delCommunityID string) error {
	fetchOwner, err := ouc.or.SelectOwnerByUserIDAndCommunityID(userID, delCommunityID)
	if err != nil {
		return err
	}
	if fetchOwner.Role != domain.Admin.String() {
		return fmt.Errorf("Only the admin is authorized to delete")
	}

	return ouc.or.DeleteOwner(delUserID, delCommunityID)
}
