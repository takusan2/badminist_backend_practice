package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/takuya-okada-01/badminist-backend/config"
	"github.com/takuya-okada-01/badminist-backend/controller"
	"github.com/takuya-okada-01/badminist-backend/infrastructure/database"
	user_repository "github.com/takuya-okada-01/badminist-backend/infrastructure/repository/user"
	auth_usecase "github.com/takuya-okada-01/badminist-backend/usecase/auth_usecase"
)

func main() {
	godotenv.Load(config.ProjectRootPath + "/.env")

	db := database.Connect()
	defer fmt.Print("db closed\n")

	router := gin.Default()
	store := cookie.NewStore([]byte(os.Getenv("SECRET_KEY")))
	router.Use(sessions.Sessions("AccessToken", store))

	ur := user_repository.NewUserRepository(db)

	au := auth_usecase.NewAuthUseCase(ur)
	v1 := router.Group("/")
	controller.NewAuthController(v1, au)

	router.Run()

}
