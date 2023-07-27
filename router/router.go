package router

import (
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/takuya-okada-01/badminist-backend/controller"
)

func NewRouter(
	ac controller.AuthController,
	uc controller.UserController,
	cc controller.CommunityController,
	oc controller.OwnerController,
	pc controller.PlayerController,
) *echo.Echo {
	e := echo.New()
	e.POST("/signup", ac.SignUpWithEmailAndPassword)
	e.POST("/login", ac.LoginWithEmailAndPassword)
	e.POST("/logout", ac.Logout)

	ug := e.Group("/users")
	ug.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET_KEY")),
		TokenLookup: "cookie:token",
	}))
	ug.GET("", uc.SelectUser)
	ug.PUT("/:id", uc.UpdateUser)

	cg := e.Group("/communities")
	cg.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET_KEY")),
		TokenLookup: "cookie:token",
	}))
	cg.POST("", cc.InsertCommunity)
	cg.GET("/:id", cc.SelectCommunityByID)
	cg.GET("", cc.SelectCommunitiesByUserID)
	cg.DELETE("/:id", cc.DeleteCommunity)

	og := e.Group("/owners")
	og.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET_KEY")),
		TokenLookup: "cookie:token",
	}))
	og.POST("", oc.InsertOwner)
	og.GET("/:id", oc.SelectOwnersByCommunityID)
	og.DELETE("", oc.DeleteOwner)

	pg := e.Group("/players")
	pg.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET_KEY")),
		TokenLookup: "cookie:token",
	}))
	pg.POST("", pc.InsertPlayer)
	pg.GET("/:community-id", pc.SelectPlayersByCommunityID)
	pg.GET("/attend/:community-id", pc.SelectAttendPlayers)
	pg.DELETE("/:id", pc.DeletePlayer)

	return e
}
