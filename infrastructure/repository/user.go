package repository

import (
	"github.com/takuya-okada-01/badminist-backend/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.IUserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) InsertUser(user *domain.User) (string, error) {
	var err error
	if err != nil {
		return "", err
	}

	result := u.db.Table("users").Create(user)
	err = result.Error
	if err != nil {
		return "", err
	}
	return user.ID, err
}

func (u *userRepository) SelectUser(userID string) (domain.User, error) {
	user := domain.User{}
	err := u.db.Select("*").Where("id = ?", userID).First(&user).Error
	return user, err
}

func (u *userRepository) UpdateUser(user *domain.User) error {
	err := u.db.Model(&user).Where("id = ?", user.ID).Updates(
		map[string]any{
			"name": user.Name,
		},
	).Error
	return err
}

func (u *userRepository) DeleteUser(id string) error {
	err := u.db.Where("id = ?", id).Delete(&domain.User{}).Error
	return err
}

func (u *userRepository) SelectUserByEmail(email string) (domain.User, error) {
	var user domain.User
	err := u.db.Select("*").Where("email = ?", email).First(&user).Error
	return user, err
}
