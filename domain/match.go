package domain

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Match struct {
	ID          int64     `gorm:"type:bigint;primary_key;auto_increment;"`
	IsSingles   bool      `gorm:"type:boolean;not null;"`
	PlayerID1   string    `gorm:"type:varchar(36);not null;"`
	PlayerID2   string    `gorm:"type:varchar(36);not null;"`
	PlayerID3   string    `gorm:"type:varchar(36);not null;"`
	PlayerID4   string    `gorm:"type:varchar(36);not null;"`
	CommunityID string    `gorm:"type:varchar(36);not null;"`
	Community   Community `gorm:"foreignKey:CommunityID;references:ID"`
	CreatedAt   time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;"`
}

type MatchCriteria struct {
	ID                   int64
	IDIsNotNull          bool
	IsSingles            bool
	IsSinglesIsNotNull   bool
	PlayerID1            string
	PlayerID1IsNotNull   bool
	PlayerID2            string
	PlayerID2IsNotNull   bool
	PlayerID3            string
	PlayerID3IsNotNull   bool
	PlayerID4            string
	PlayerID4IsNotNull   bool
	CommunityID          string
	CommunityIDIsNotNull bool
}

type IMatchRepository interface {
	InsertMatch(communityID string, match *Match) (int64, error)
	SelectMatch(criteria MatchCriteria) (Match, error)
	SelectMatches(criteria MatchCriteria) ([]Match, error)
	UpdateMatch(match *Match) error
	DeleteMatch(id int64) error
}

type IMatchUseCase interface {
	SelectMatch(ctx *gin.Context, communityID string, id string) (Match, error)
	SelectMatchesByCommunityID(ctx *gin.Context, communityId string) ([]Match, error)
	SelectMatchesByCommunityIDAndDate(ctx *gin.Context, communityId string, date string) ([]Match, error)
	SelectMatchesByPlayerIDAndDate(ctx *gin.Context, playerID string, date string) ([]Match, error)
	UpdateMatch(ctx *gin.Context, communityID string, id int64, isSingles bool, playerID1 string, playerID2 string, playerID3 string, playerID4 string) error
	DeleteMatch(ctx *gin.Context, communityID string, id int64) error
}
