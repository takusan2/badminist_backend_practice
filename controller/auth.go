package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/takuya-okada-01/badminist-backend/domain"
)

type AuthController interface {
	SignUpWithEmailAndPassword(ctx echo.Context) error
	LoginWithEmailAndPassword(ctx echo.Context) error
	Logout(ctx echo.Context) error
	RefreshToken(ctx echo.Context) error
}
type authController struct {
	au domain.IAuthUseCase
}

func NewAuthController(au domain.IAuthUseCase) AuthController {
	return &authController{au: au}
}

func (ac *authController) SignUpWithEmailAndPassword(ctx echo.Context) error {
	user := domain.User{}
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	_, err := ac.au.SignUpWithEmailAndPassword(user.Email, user.Password)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, user)
}

func (ac *authController) LoginWithEmailAndPassword(ctx echo.Context) error {
	user := domain.User{}
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	tokenString, err := ac.au.LoginWithEmailAndPassword(user.Email, user.Password)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	ctx.SetCookie(&http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Path:     "/",
		Domain:   os.Getenv("API_DOMAIN"),
		Expires:  time.Now().Add(24 * time.Hour),
		MaxAge:   60 * 60 * 24,
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	return ctx.JSON(http.StatusOK, user)
}

func (ac *authController) Logout(ctx echo.Context) error {
	ctx.SetCookie(&http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		Domain:   os.Getenv("API_DOMAIN"),
		Expires:  time.Now(),
		MaxAge:   -1,
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
	return ctx.NoContent(http.StatusOK)
}

func (ac *authController) RefreshToken(ctx echo.Context) error {
	return nil
}

func GetCurrentUser(ctx echo.Context) string {
	user := ctx.Get("user").(*jwt.Token)
	userClaims := user.Claims.(jwt.MapClaims)
	userID := userClaims["user_id"].(string)
	return userID
}
