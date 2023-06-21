package repository

import (
	"github.com/takuya-okada-01/badminist-backend/domain"
	"gorm.io/gorm"
)

type playerRepository struct {
	db *gorm.DB
}

func NewPlayerRepository(db *gorm.DB) domain.IPlayerRepository {
	return &playerRepository{db: db}
}

func (p *playerRepository) InsertPlayer(player *domain.Player) (string, error) {
	var err error

	result := p.db.Table("players").Create(player)
	err = result.Error
	if err != nil {
		return "", err
	}
	return player.ID, err
}

func (p *playerRepository) SelectPlayer(criteria domain.PlayerCriteria) (domain.Player, error) {
	var player domain.Player
	err := p.db.Select("*").Where(criteria).First(&player).Error
	return player, err
}

func (p *playerRepository) SelectPlayers(criteria domain.PlayerCriteria) ([]domain.Player, error) {
	var players []domain.Player
	err := p.db.Select("*").Where(criteria).Find(&players).Error
	return players, err
}

func (p *playerRepository) UpdatePlayer(player *domain.Player) error {
	err := p.db.Model(&player).Where("id = ?", player.ID).Updates(player).Error
	return err
}

func (p *playerRepository) DeletePlayer(id string) error {
	err := p.db.Where("id = ?", id).Delete(&domain.Player{}).Error
	return err
}
