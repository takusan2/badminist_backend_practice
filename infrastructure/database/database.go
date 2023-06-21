package database

import (
	"fmt"
	"os"

	sql "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
	"github.com/takuya-okada-01/badminist-backend/config"
	"github.com/takuya-okada-01/badminist-backend/domain"
)

func Connect() *gorm.DB {
	godotenv.Load(config.ProjectRootPath + "/.env")
	cfg := sql.Config{
		User:      os.Getenv("DB_USER"),
		Passwd:    os.Getenv("DB_PASS"),
		Net:       os.Getenv("DB_NET"),
		Addr:      os.Getenv("DB_ADDR"),
		DBName:    os.Getenv("DB_NAME"),
		ParseTime: true,
	}
	dsn := cfg.FormatDSN()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print("failed to Connected to database\n")
		panic(err.Error())
	}
	db.Logger.LogMode(3)
	db.AutoMigrate(&domain.User{}, &domain.Community{}, &domain.Player{}, &domain.Owner{}, &domain.Match{})
	fmt.Print("Connected to database!!\n")
	return db
}
