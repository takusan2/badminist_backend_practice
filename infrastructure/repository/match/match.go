package repository

import (
	"github.com/takuya-okada-01/badminist-backend/domain"
	"github.com/takuya-okada-01/badminist-backend/utils"
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

func (m *matchRepository) SelectMatch(criteria domain.MatchCriteria) (domain.Match, error) {
	var match domain.Match
	mapCriteria := utils.CriteriaToMap(criteria)
	err := m.db.Select("*").Where(mapCriteria).First(&match).Error
	return match, err
}

func (m *matchRepository) SelectMatches(criteria domain.MatchCriteria) ([]domain.Match, error) {
	var matches []domain.Match
	mapCriteria := utils.CriteriaToMap(criteria)
	err := m.db.Select("*").Where(mapCriteria).Find(&matches).Error
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
