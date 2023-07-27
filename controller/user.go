package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/takuya-okada-01/badminist-backend/domain"
	"github.com/takuya-okada-01/badminist-backend/usecase"
)

type UserController interface {
	SelectUser(ctx echo.Context) error
	UpdateUser(ctx echo.Context) error
	DeleteUser(ctx echo.Context) error
}
type userController struct {
	uu usecase.UserUseCase
}

func NewUserController(uu usecase.UserUseCase) UserController {
	return &userController{uu: uu}
}

func (uc *userController) SelectUser(ctx echo.Context) error {
	userID := GetCurrentUser(ctx)
	resUser, err := uc.uu.SelectUser(userID)
	if err != nil {
		return ctx.JSON(500, err.Error())
	}
	return ctx.JSON(200, resUser)
}

func (uc *userController) UpdateUser(ctx echo.Context) error {
	userID := GetCurrentUser(ctx)
	paramUserID := ctx.Param("id")
	if userID != paramUserID {
		return ctx.JSON(400, map[string]string{"message": "invalid user id"})
	}

	user := domain.User{}
	ctx.Bind(&user)
	resUser, err := uc.uu.UpdateUser(
		userID,
		user,
	)
	if err != nil {
		return ctx.JSON(500, err.Error())
	}
	return ctx.JSON(200, resUser)
}

func (uc *userController) DeleteUser(ctx echo.Context) error {
	userID := GetCurrentUser(ctx)
	err := uc.uu.DeleteUser(userID)
	if err != nil {
		return ctx.JSON(500, err.Error())
	}
	return ctx.JSON(200, map[string]string{"message": "success"})
}
