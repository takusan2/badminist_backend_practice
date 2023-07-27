package usecase

import (
	"encoding/base64"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/takuya-okada-01/badminist-backend/domain"
	"golang.org/x/crypto/bcrypt"
)

type authUseCase struct {
	userRepository domain.IUserRepository
}

func NewAuthUseCase(userRepository domain.IUserRepository) domain.IAuthUseCase {
	return &authUseCase{userRepository: userRepository}
}

func (au *authUseCase) SignUpWithEmailAndPassword(email string, password string) (string, error) {
	salt := base64.StdEncoding.EncodeToString(uuid.New().NodeID())
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user := domain.User{
		Email:    email,
		Salt:     salt,
		Password: string(passwordHash),
	}

	id, err := au.userRepository.InsertUser(&user)
	if err != nil {
		return "", err
	}
	return id, err
}

func (au *authUseCase) LoginWithEmailAndPassword(email string, password string) (string, error) {
	user, err := au.userRepository.SelectUserByEmail(email)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password+user.Salt))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, err
}
