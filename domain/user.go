package domain

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TokenString = string

type User struct {
	ID           string    `gorm:"type:varchar(36);primary_key;"`
	Name         string    `gorm:"type:varchar(255);not null;"`
	Email        string    `gorm:"type:varchar(255);not null;unique;"`
	PasswordHash string    `gorm:"type:varchar(255);not null;"`
	Salt         string    `gorm:"type:varchar(255);not null;"`
	CreatedAt    time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;"`
	UpdatedAt    time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return nil
}

type UserCriteria struct {
	ID             string
	IDIsNotNull    bool
	Name           string
	NameIsNotNull  bool
	Email          string
	EmailIsNotNull bool
}

type IUserRepository interface {
	InsertUser(user *User) (string, error)
	SelectUser(criteria UserCriteria) (User, error)
	UpdateUser(user *User) error
	DeleteUser(id string) error
}

type IUserUseCase interface {
	UpdateUser(ctx *gin.Context, id string, name string) error
	DeleteUser(ctx *gin.Context, id string) error
}

type IAuthUseCase interface {
	SignUpWithEmailAndPassword(ctx *gin.Context, email string, password string) (TokenString, error)
	LoginWithEmailAndPassword(ctx *gin.Context, email string, password string) (TokenString, error)
	Logout(ctx *gin.Context) error
}
