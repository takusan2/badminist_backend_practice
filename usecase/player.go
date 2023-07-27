package usecase

import (
	"fmt"

	"github.com/takuya-okada-01/badminist-backend/domain"
)

type PlayerUseCase interface {
	InsertPlayer(userID string, player *domain.Player) (string, error)
	SelectPlayersByCommunityID(communityID string) ([]domain.Player, error)
	SelectAttendPlayers(communityID string) ([]domain.Player, error)
	UpdatePlayer(player *domain.Player) error
	DeletePlayer(userID string, id string) error
}

type playerUseCase struct {
	pr domain.IPlayerRepository
	or domain.IOwnerRepository
}

func NewPlayerUseCase(
	pr domain.IPlayerRepository,
	or domain.IOwnerRepository,
) PlayerUseCase {
	return &playerUseCase{
		pr: pr,
		or: or,
	}
}

func (p *playerUseCase) InsertPlayer(userID string, player *domain.Player) (string, error) {
	owner, err := p.or.SelectOwnerByUserIDAndCommunityID(userID, player.CommunityID)
	if err != nil {
		return "", err
	}
	if owner.Role != domain.Admin.String() {
		return "", fmt.Errorf("Only the admin is authorized to insert")
	}
	return p.pr.InsertPlayer(player)
}

func (p *playerUseCase) SelectPlayersByCommunityID(communityID string) ([]domain.Player, error) {
	return p.pr.SelectPlayersByCommunityID(communityID)
}

func (p *playerUseCase) SelectAttendPlayers(communityID string) ([]domain.Player, error) {

	return p.pr.SelectAttendPlayers(communityID)
}

func (p *playerUseCase) UpdatePlayer(player *domain.Player) error {
	return p.pr.UpdatePlayer(player)
}

func (p *playerUseCase) DeletePlayer(userID string, id string) error {
	owner, err := p.or.SelectOwnerByUserIDAndCommunityID(userID, id)
	if err != nil {
		return err
	}
	if owner.Role != domain.Admin.String() {
		return fmt.Errorf("Only the admin is authorized to delete")
	}
	return p.pr.DeletePlayer(id)
}
