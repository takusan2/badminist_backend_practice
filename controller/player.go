package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/takuya-okada-01/badminist-backend/domain"
	"github.com/takuya-okada-01/badminist-backend/usecase"
)

type PlayerController interface {
	InsertPlayer(ctx echo.Context) error
	SelectPlayersByCommunityID(ctx echo.Context) error
	SelectAttendPlayers(ctx echo.Context) error
	DeletePlayer(ctx echo.Context) error
}

type playerController struct {
	pu usecase.PlayerUseCase
}

func NewPlayerController(pu usecase.PlayerUseCase) PlayerController {
	return &playerController{pu: pu}
}

func (pc *playerController) InsertPlayer(ctx echo.Context) error {
	userID := GetCurrentUser(ctx)
	player := domain.Player{}
	ctx.Bind(&player)
	id, err := pc.pu.InsertPlayer(userID, &player)
	if err != nil {
		return ctx.JSON(500, err.Error())
	}
	return ctx.JSON(200, map[string]string{"id": id})
}

func (pc *playerController) SelectPlayersByCommunityID(ctx echo.Context) error {
	communityID := ctx.Param("community-id")
	players, err := pc.pu.SelectPlayersByCommunityID(communityID)
	if err != nil {
		return ctx.JSON(500, err.Error())
	}
	return ctx.JSON(200, players)
}

func (pc *playerController) SelectAttendPlayers(ctx echo.Context) error {
	communityID := ctx.Param("community-id")
	players, err := pc.pu.SelectAttendPlayers(communityID)
	if err != nil {
		return ctx.JSON(500, err.Error())
	}
	return ctx.JSON(200, players)
}

func (pc *playerController) DeletePlayer(ctx echo.Context) error {
	user := GetCurrentUser(ctx)
	id := ctx.Param("id")
	err := pc.pu.DeletePlayer(user, id)
	if err != nil {
		return ctx.JSON(500, err.Error())
	}
	return ctx.JSON(200, map[string]string{"id": id})
}
