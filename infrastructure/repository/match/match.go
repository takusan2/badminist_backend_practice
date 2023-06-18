package repository

import (
	"github.com/takuya-okada-01/badminist-backend/domain"
	"gorm.io/gorm"
)

type matchRepository struct {
	db *gorm.DB
}

func NewMatchRepository(db *gorm.DB) domain.IMatchRepository {
	return &matchRepository{db: db}
}

func (m *matchRepository) InsertMatch(communityID string, match *domain.Match) (int64, error) {
	var err error
	match.CommunityID = communityID

	result := m.db.Table("matches").Create(match)
	err = result.Error
	if err != nil {
		return 0, err
	}
	return match.ID, err
}

func (m *matchRepository) SelectMatch(id int64) (domain.Match, error) {
	var match domain.Match
	err := m.db.Select("*").Where("id = ?", id).First(&match).Error
	return match, err
}

func (m *matchRepository) SelectMatchesByCommunityID(communityId string) ([]domain.Match, error) {
	var matches []domain.Match
	err := m.db.Select("*").Where("community_id = ?", communityId).Find(&matches).Error
	return matches, err
}

func (m *matchRepository) SelectMatchesByCommunityIDAndDate(communityId string, date string) ([]domain.Match, error) {
	var matches []domain.Match
	err := m.db.Select("*").Where(map[string]interface{}{"community_id": communityId, "date": date}).Find(&matches).Error
	return matches, err
}

func (m *matchRepository) SelectMatchesByPlayerIDAndDate(playerId string, date string) ([]domain.Match, error) {
	var matches []domain.Match
	err := m.db.Select("*").Where(map[string]interface{}{"player_id": playerId, "date": date}).Find(&matches).Error
	return matches, err
}

func (m *matchRepository) UpdateMatch(match *domain.Match) error {
	err := m.db.Model(&match).Where("id = ?", match.ID).Updates(match).Error
	return err
}

func (m *matchRepository) DeleteMatch(id int64) error {
	err := m.db.Where("id = ?", id).Delete(&domain.Match{}).Error
	return err
}
