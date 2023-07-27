package usecase

import (
	"github.com/takuya-okada-01/badminist-backend/domain"
	"github.com/takuya-okada-01/badminist-backend/validator"
)

type UserUseCase interface {
	SelectUser(id string) (domain.UserResopnse, error)
	UpdateUser(id string, user domain.User) (domain.UserResopnse, error)
	DeleteUser(id string) error
}

type userUseCase struct {
	ur domain.IUserRepository
	uv validator.UserValidator
}

func NewUserUseCase(
	ur domain.IUserRepository,
	uv validator.UserValidator,
) UserUseCase {
	return &userUseCase{ur: ur, uv: uv}
}

func (u *userUseCase) SelectUser(id string) (domain.UserResopnse, error) {
	user, err := u.ur.SelectUser(id)
	if err != nil {
		return domain.UserResopnse{}, err
	}

	return domain.UserResopnse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}

func (u *userUseCase) UpdateUser(id string, user domain.User) (domain.UserResopnse, error) {
	_, err := u.ur.SelectUser(id)
	if err != nil {
		return domain.UserResopnse{}, err
	}
	if err = u.ur.UpdateUser(&user); err != nil {
		return domain.UserResopnse{}, err
	}

	return domain.UserResopnse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}

func (u *userUseCase) DeleteUser(id string) error {
	return u.ur.DeleteUser(id)
}
