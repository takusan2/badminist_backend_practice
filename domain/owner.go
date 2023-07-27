package domain

type Owner struct {
	UserID      string    `gorm:"primary_key;"`
	User        User      `gorm:"foreignKey:UserID;references:ID"`
	CommunityID string    `gorm:"primary_key;"`
	Community   Community `gorm:"foreignKey:CommunityID;references:ID"`
	Role        string    `gorm:"type:varchar(255);not null;"`
}
type IOwnerRepository interface {
	InsertOwner(owner Owner) error
	SelectOwnerByUserIDAndCommunityID(userID string, communityID string) (Owner, error)
	SelectOwnersByCommunityID(communityID string) ([]Owner, error)
	SelectOwnersByUserID(userID string) ([]Owner, error)
	UpdateOwner(owner *Owner) error
	DeleteOwner(userID string, communityID string) error
}

type Role int

const (
	Admin Role = iota
	Staff
	Member
)

func (r Role) String() string {
	switch r {
	case Admin:
		return "admin"
	case Staff:
		return "staff"
	case Member:
		return "member"
	default:
		return "unknown"
	}
}
