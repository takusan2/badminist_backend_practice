package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/takuya-okada-01/badminist-backend/domain"
	"github.com/takuya-okada-01/badminist-backend/usecase"
)

type CommunityController interface {
	InsertCommunity(ctx echo.Context) error
	SelectCommunityByID(ctx echo.Context) error
	SelectCommunitiesByUserID(ctx echo.Context) error
	DeleteCommunity(ctx echo.Context) error
}

type communityController struct {
	cu usecase.CommunityUseCase
}

func NewCommunityController(cu usecase.CommunityUseCase) CommunityController {
	return &communityController{cu: cu}
}

func (cc *communityController) InsertCommunity(ctx echo.Context) error {
	userID := GetCurrentUser(ctx)
	community := domain.Community{}
	ctx.Bind(&community)
	id, err := cc.cu.InsertCommunity(userID, community)
	if err != nil {
		return ctx.JSON(500, err.Error())
	}
	return ctx.JSON(200, map[string]string{"id": id})
}

func (cc *communityController) SelectCommunityByID(ctx echo.Context) error {
	userID := GetCurrentUser(ctx)
	id := ctx.Param("id")
	resCommunity, err := cc.cu.SelectCommunityByID(userID, id)
	if err != nil {
		return ctx.JSON(500, err.Error())
	}
	return ctx.JSON(200, resCommunity)
}

func (cc *communityController) SelectCommunitiesByUserID(ctx echo.Context) error {
	userID := GetCurrentUser(ctx)
	resCommunities, err := cc.cu.SelectCommunitiesByUserID(userID)
	if err != nil {
		return ctx.JSON(500, err.Error())
	}
	return ctx.JSON(200, resCommunities)
}

func (cc *communityController) DeleteCommunity(ctx echo.Context) error {
	userID := GetCurrentUser(ctx)
	id := ctx.Param("id")
	err := cc.cu.DeleteCommunity(userID, id)
	if err != nil {
		return ctx.JSON(500, err.Error())
	}
	return ctx.JSON(200, map[string]string{"message": "success"})
}
