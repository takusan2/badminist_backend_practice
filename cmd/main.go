package main

import (
	"github.com/takuya-okada-01/badminist-backend/controller"
	"github.com/takuya-okada-01/badminist-backend/infrastructure/database"
	"github.com/takuya-okada-01/badminist-backend/infrastructure/repository"
	"github.com/takuya-okada-01/badminist-backend/router"
	"github.com/takuya-okada-01/badminist-backend/usecase"
	"github.com/takuya-okada-01/badminist-backend/validator"
)

func main() {
	db := database.Connect()
	defer database.CloseDB(db)

	ur := repository.NewUserRepository(db)
	cr := repository.NewCommunityRepository(db)
	or := repository.NewOwnerRepository(db)
	pr := repository.NewPlayerRepository(db)

	uv := validator.NewUserValidator()

	uu := usecase.NewUserUseCase(ur, uv)
	au := usecase.NewAuthUseCase(ur)
	cu := usecase.NewCommunityUseCase(cr, or)
	ou := usecase.NewOwnerUseCase(or, cr)
	pu := usecase.NewPlayerUseCase(pr, or)

	ac := controller.NewAuthController(au)
	uc := controller.NewUserController(uu)
	cc := controller.NewCommunityController(cu)
	oc := controller.NewOwnerController(ou)
	pc := controller.NewPlayerController(pu)

	e := router.NewRouter(ac, uc, cc, oc, pc)
	e.Logger.Fatal(e.Start(":8080"))
}
