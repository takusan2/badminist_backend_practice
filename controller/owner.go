package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/takuya-okada-01/badminist-backend/domain"
	"github.com/takuya-okada-01/badminist-backend/usecase"
)

type OwnerController interface {
	InsertOwner(ctx echo.Context) error
	SelectOwnersByCommunityID(ctx echo.Context) error
	UpdateOwner(ctx echo.Context) error
	DeleteOwner(ctx echo.Context) error
}

type ownerController struct {
	ou usecase.OwnerUseCase
}

func NewOwnerController(ou usecase.OwnerUseCase) OwnerController {
	return &ownerController{ou: ou}
}

func (oc *ownerController) InsertOwner(ctx echo.Context) error {
	userID := GetCurrentUser(ctx)
	owner := domain.Owner{}
	if err := ctx.Bind(&owner); err != nil {
		return ctx.JSON(400, map[string]string{"message": "invalid request"})
	}

	err := oc.ou.InsertOwner(userID, owner)
	if err != nil {
		return ctx.JSON(500, err.Error())
	}
	return ctx.JSON(200, map[string]string{"message": "success"})
}

func (oc *ownerController) SelectOwnersByCommunityID(ctx echo.Context) error {
	userID := GetCurrentUser(ctx)
	communityID := ctx.Param("id")
	resOwners, err := oc.ou.SelectOwnersByCommunityID(userID, communityID)
	if err != nil {
		return ctx.JSON(500, err.Error())
	}
	return ctx.JSON(200, resOwners)
}

func (oc *ownerController) UpdateOwner(ctx echo.Context) error {
	userID := GetCurrentUser(ctx)
	owner := domain.Owner{}
	if err := ctx.Bind(&owner); err != nil {
		return ctx.JSON(400, map[string]string{"message": "invalid request"})
	}

	err := oc.ou.UpdateOwner(userID, &owner)
	if err != nil {
		return ctx.JSON(500, err.Error())
	}
	return ctx.JSON(200, map[string]string{"message": "success"})
}

func (oc *ownerController) DeleteOwner(ctx echo.Context) error {
	userID := GetCurrentUser(ctx)
	owner := domain.Owner{}
	if err := ctx.Bind(&owner); err != nil {
		return ctx.JSON(400, map[string]string{"message": "invalid request"})
	}
	err := oc.ou.DeleteOwner(userID, owner.UserID, owner.CommunityID)
	if err != nil {
		return ctx.JSON(500, err.Error())
	}
	return ctx.JSON(200, map[string]string{"message": "success"})
}
