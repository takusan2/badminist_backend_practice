package main

import (
	"fmt"

	"github.com/takuya-okada-01/badminist-backend/domain"
	"github.com/takuya-okada-01/badminist-backend/infrastructure/database"
)

func main() {
	dbConn := database.Connect()
	defer fmt.Println("Successfully migrated")
	defer database.CloseDB(dbConn)
	dbConn.AutoMigrate(&domain.User{}, &domain.Community{}, &domain.Owner{}, &domain.Player{})
}
