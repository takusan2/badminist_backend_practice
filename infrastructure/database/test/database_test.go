package database_test

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/takuya-okada-01/badminist-backend/infrastructure/database"
)

func TestConnect(t *testing.T) {
	godotenv.Load("../../../.env")
	db := database.Connect()
	if db == nil {
		t.Errorf("Connect() == %v, want %v", db, "not nil")
	}
}
