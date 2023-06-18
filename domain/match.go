package domain

import "time"

type Match struct {
	ID          string    `gorm:"type:varchar(36);primary_key;"`
	IsSingles   bool      `gorm:"type:boolean;not null;"`
	PlayerID1   string    `gorm:"type:varchar(36);not null;"`
	PlayerID2   string    `gorm:"type:varchar(36);not null;"`
	PlayerID3   string    `gorm:"type:varchar(36);not null;"`
	PlayerID4   string    `gorm:"type:varchar(36);not null;"`
	CommunityID string    `gorm:"type:varchar(36);not null;"`
	Community   Community `gorm:"foreignKey:CommunityID;references:ID"`
	CreatedAt   time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;"`
}

type IMatchRepository interface {
	InsertMatch(communityID string, match *Match) (string, error)
	SelectMatch(communityID string, id string) (Match, error)
}
