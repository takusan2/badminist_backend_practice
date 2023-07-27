package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TokenString = string

type User struct {
	ID        string    `json:"id" gorm:"type:varchar(36);primary_key;"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null;"`
	Email     string    `json:"email" gorm:"type:varchar(255);not null;unique;"`
	Password  string    `json:"password" gorm:"type:varchar(255);not null;"`
	Salt      string    `json:"salt" gorm:"type:varchar(255);not null;"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return nil
}

type UserResopnse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type IUserRepository interface {
	InsertUser(user *User) (string, error)
	SelectUser(id string) (User, error)
	SelectUserByEmail(mail string) (User, error)
	UpdateUser(user *User) error
	DeleteUser(id string) error
}

type IAuthUseCase interface {
	SignUpWithEmailAndPassword(email string, password string) (TokenString, error)
	LoginWithEmailAndPassword(email string, password string) (TokenString, error)
}
